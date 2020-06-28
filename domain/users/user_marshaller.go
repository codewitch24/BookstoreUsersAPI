package users

import "encoding/json"

type PublicUser struct {
	Id      int64  `json:"id"`
	Created string `json:"created"`
	Status  string `json:"status"`
}

type PrivateUser struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Created   string `json:"created"`
	Status    string `json:"status"`
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for i, u := range users {
		result[i] = u.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:      user.Id,
			Created: user.Created,
			Status:  user.Status,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	if err := json.Unmarshal(userJson, &privateUser); err != nil {
		return nil
	}
	return privateUser
}
