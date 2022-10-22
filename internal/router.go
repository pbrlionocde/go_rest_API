package internal

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	Endpoints(router)
	router.Run()
}
