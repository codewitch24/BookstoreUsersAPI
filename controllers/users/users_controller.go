package users

import (
	"github.com/codewitch24/BookstoreUsersAPI/domain/users"
	"github.com/codewitch24/BookstoreUsersAPI/services"
	"github.com/codewitch24/BookstoreUsersAPI/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		e := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(e.Status, e)
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
	userId, userErr := strconv.ParseInt(c.Param("UserId"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("User Id Should be a Number")
		c.JSON(err.Status, err)
		return
	}
	result, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
