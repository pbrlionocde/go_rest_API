package service

import (
	"f_gin/pkg/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, storage.Storage)
}
