package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovery: panic recovered: %s", err)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
