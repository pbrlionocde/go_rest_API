package internal

import (
	"f_gin/pkg/service/handler/auth"
	"f_gin/pkg/service/handler/middleware"

	"github.com/gin-gonic/gin"
)

func Endpoints(router *gin.Engine) {
	publicGroup(router)
	privateGroup(router)
}


func publicGroup(router *gin.Engine) {
	public_group := router.Group("/public")
	public_group.POST("/register", auth.CreateUser)
	public_group.POST("/get_token", auth.GetToken)
}

func privateGroup(router *gin.Engine) {
	private_group := router.Group("/api")
	private_group.Use(middleware.JWTAuthMiddleware())
    private_group.GET("/test", auth.TestToken)
}

