package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Authenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("Authenticated okok")
	}
}
