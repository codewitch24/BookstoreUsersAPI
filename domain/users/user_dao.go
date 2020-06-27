package users

import (
	"github.com/codewitch24/BookstoreUsersAPI/datasources/mysql/usersdb"
	"github.com/codewitch24/BookstoreUsersAPI/utils/date"
	"github.com/codewitch24/BookstoreUsersAPI/utils/errors"
	"github.com/codewitch24/BookstoreUsersAPI/utils/mysql"
)

const (
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, created) VALUES (?, ?, ?, ?);"
	queryGetUser    = "SELECT * FROM users WHERE id = ?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?"
)

func (user *User) Get() *errors.RestError {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
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
		return mysql.ParseError(err)
	}
	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer func() {
		_ = stmt.Close()
	}()
	user.Created = date.GetNowStringDate()
	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Created)
	if err != nil {
		return mysql.ParseError(err)
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return mysql.ParseError(err)
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestError {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer func() {
		_ = stmt.Close()
	}()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql.ParseError(err)
	}
	return nil
}

func (user *User) Delete() (int64, *errors.RestError) {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return 0, errors.NewInternalServerError(err.Error())
	}
	defer func() {
		_ = stmt.Close()
	}()
	result, err := stmt.Exec(user.Id)
	if err != nil {
		return 0, mysql.ParseError(err)
	}
	n, err := result.RowsAffected()
	if err != nil {
		return 0, errors.NewInternalServerError(err.Error())
	}
	return n, nil
}
