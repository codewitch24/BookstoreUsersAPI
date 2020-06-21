package app

import (
	PingController "github.com/codewitch24/BookstoreUsersAPI/controllers/ping"
	UsersController "github.com/codewitch24/BookstoreUsersAPI/controllers/users"
	"github.com/gin-gonic/gin"
)

func MapUrls(r *gin.Engine) *gin.Engine {
	r.GET("/ping", PingController.Ping)
	r.POST("/users", UsersController.CreateUser)
	r.GET("/users/:UserId", UsersController.GetUser)
	return r
}
