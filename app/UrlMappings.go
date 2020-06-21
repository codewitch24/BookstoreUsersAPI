package app

import (
	"github.com/codewitch24/BookstoreUsersAPI/controllers"
	"github.com/gin-gonic/gin"
)

func MapUrls(r *gin.Engine) *gin.Engine {
	r.GET("/ping", controllers.Ping)
	return r
}
