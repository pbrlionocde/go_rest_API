package internal

import (
	"f_gin/pkg/service"

	"github.com/gin-gonic/gin"
)

func Endpoints(router *gin.Engine) {
	router.GET("/cars", service.Cars)
}
