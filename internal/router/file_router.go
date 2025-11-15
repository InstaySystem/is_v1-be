package router

import (
	"github.com/InstaySystem/is-be/internal/handler"
	"github.com/gin-gonic/gin"
)

func FileRouter(rg *gin.RouterGroup, hdl *handler.FileHandler) {
	file := rg.Group("/files")
	{
		file.POST("/presigned-urls/uploads", hdl.UploadPresignedURLs)

		file.POST("/presigned-urls/views", hdl.ViewPresignedURLs)
	}
}
