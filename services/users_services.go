package services

import (
	"github.com/codewitch24/BookstoreUsersAPI/domain/users"
	"github.com/codewitch24/BookstoreUsersAPI/utils/date"
	"github.com/codewitch24/BookstoreUsersAPI/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusNonActive
	user.Created = date.GetNowDatabaseFormat()
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.FirstName != "" {
			current.LastName = user.LastName
		}
		if user.FirstName != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userId int64) *errors.RestError {
	user := &users.User{Id: userId}
	result, err := user.Delete()
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.NewBadRequestError("Maybe invalid data to delete")
	}
	return nil
}

func Search(status string) ([]users.User, *errors.RestError) {
	user := &users.User{}
	return user.FindByStatus(status)
}
