package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Установка заголовков CORS
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, PATCH, GET, PUT, DELETE")

		// Обработка preflight-запросов (OPTIONS)
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// Переход к следующему обработчику
		c.Next()
	}
}
