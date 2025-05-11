package router

import (
	"go-send-email/handlers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func checkEnvMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requiredEnv := []string{"CONFIG_AUTH_EMAIL", "CONFIG_AUTH_PASSWORD", "CONFIG_SENDER_NAME"}
		for _, key := range requiredEnv {
			if os.Getenv(key) == "" {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Missing required environment variables",
				})
				return
			}
		}
		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(checkEnvMiddleware())

	r.POST("/send-email", handlers.SendEmailHandler)

	return r
}
