package router

import (
	"github.com/InstaySystem/is-be/internal/handler"
	"github.com/gin-gonic/gin"
)

func FileRouter(rg *gin.RouterGroup, hdl *handler.FileHandler) {
	file := rg.Group("/files")
	{
		file.POST("/presigned-url", hdl.CreatePresignedURL)

		file.GET("/view", hdl.ViewFile)
	}
}
