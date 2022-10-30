package middleware

import (
	"f_gin/pkg/service/handler/auth"
	"f_gin/pkg/service/logger"
	// "fmt"
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

var errorLogger *log.Logger

func init() {
	errorLogger = logger.GetAuthorizeErrorLogger()
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatus(401)
		}

		tokenUser, err := auth.ParseToken(token)
		if err!= nil {
			errorLogger.Println(err)
            c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("tokenUser", tokenUser)
	}
}
