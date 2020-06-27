package app

import (
	"github.com/codewitch24/BookstoreUsersAPI/controllers/ping"
	"github.com/codewitch24/BookstoreUsersAPI/controllers/users"
	"github.com/gin-gonic/gin"
)

func MapUrls(r *gin.Engine) {
	r.GET("/ping", ping.Ping)
	r.POST("/users", users.Create)
	r.GET("/users/:user_id", users.Get)
	r.PUT("/users/:user_id", users.Update)
	r.PATCH("/users/:user_id", users.Update)
	r.DELETE("/users/:user_id", users.Delete)
}
