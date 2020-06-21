package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func StartApplication() {
	r := MapUrls(router)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
