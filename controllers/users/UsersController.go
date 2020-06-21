package users

import (
	"github.com/codewitch24/BookstoreUsersAPI/domain/users"
	"github.com/codewitch24/BookstoreUsersAPI/services"
	"github.com/codewitch24/BookstoreUsersAPI/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(restError.Status, restError)
		return
	}
	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me!")
}
