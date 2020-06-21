package services

import (
	"github.com/codewitch24/BookstoreUsersAPI/domain/users"
	"github.com/codewitch24/BookstoreUsersAPI/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return &user, nil
}
