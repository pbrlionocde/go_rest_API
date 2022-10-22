package internal

import (
	"f_gin/pkg/service/handler/auth"

	"github.com/gin-gonic/gin"
)

func Endpoints(router *gin.Engine) {
	router.POST("/car", auth.CreateUser)
}
