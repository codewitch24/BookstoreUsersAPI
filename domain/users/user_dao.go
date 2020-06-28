package users

import (
	"fmt"
	"github.com/codewitch24/BookstoreUsersAPI/datasources/mysql/usersdb"
	"github.com/codewitch24/BookstoreUsersAPI/utils/errors"
	"github.com/codewitch24/BookstoreUsersAPI/utils/mysql"
)

const (
	queryInsertUser       = "INSERT INTO users (first_name, last_name, email, created, password, status) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, created, status FROM users WHERE id = ?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, created, status from users WHERE status=?;"
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
		&user.Created,
		&user.Status); err != nil {
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
	result, err := stmt.Exec(
		user.FirstName, user.LastName, user.Email, user.Created, user.Password, user.Status,
	)
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

func (user *User) FindByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer func() {
		_ = stmt.Close()
	}()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer func() {
		_ = rows.Close()
	}()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Created,
			&user.Status); err != nil {
			return nil, mysql.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(
			fmt.Sprintf("No users matching status %s", status),
		)
	}
	return results, nil
}
