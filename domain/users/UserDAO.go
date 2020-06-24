package users

import (
	"fmt"
	"github.com/codewitch24/BookstoreUsersAPI/utils/date"
	"github.com/codewitch24/BookstoreUsersAPI/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.Created = result.Created
	return nil
}

func (user *User) Save() *errors.RestError {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s aleady registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d aleady exists", user.Id))
	}
	user.Created = date.GetNowStringDate()
	usersDB[user.Id] = user
	return nil
}
