package users

import (
	"fmt"
	"github.com/codewitch24/BookstoreUsersAPI/datasources/mysql/usersdb"
	"github.com/codewitch24/BookstoreUsersAPI/utils/date"
	"github.com/codewitch24/BookstoreUsersAPI/utils/errors"
	"strings"
)

const (
	IndexUniqueEmail = "unique_email"
	ErrorNoRows      = "no rows in result set"
	QueryInsertUser  = "INSERT INTO users (first_name, last_name, email, created) VALUES (?, ?, ?, ?);"
	QueryGetUser     = "SELECT * FROM users WHERE id = ?;"
)

func (user *User) Get() *errors.RestError {
	stmt, err := usersdb.Client.Prepare(QueryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer func() {
		_ = stmt.Close()
	}()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Created); err != nil {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError(
				fmt.Sprintf("User %d not found", user.Id),
			)
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("Error when trying to get user %d: %s", user.Id, err.Error()),
		)
	}
	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := usersdb.Client.Prepare(QueryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer func() {
		_ = stmt.Close()
	}()
	user.Created = date.GetNowStringDate()
	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Created)
	if err != nil {
		if strings.Contains(err.Error(), IndexUniqueEmail) {
			return errors.NewBadRequestError(
				fmt.Sprintf("Email %s already exists", user.Email),
			)
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("Error when trying to save user: %s", err.Error()),
		)
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("Error when trying to save user %s", err.Error()),
		)
	}
	user.Id = userId
	return nil
}
