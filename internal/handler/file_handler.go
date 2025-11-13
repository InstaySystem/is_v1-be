package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/InstaySystem/is-be/internal/common"
	"github.com/InstaySystem/is-be/internal/service"
	"github.com/InstaySystem/is-be/internal/types"
	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	fileSvc service.FileService
}

func NewFileHandler(fileSvc service.FileService) *FileHandler {
	return &FileHandler{fileSvc}
}

func (h *FileHandler) CreatePresignedURL(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var req types.PresignedURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		mess := common.HandleValidationError(err)
		common.ToAPIResponse(c, http.StatusBadRequest, mess, nil)
		return
	}

	presignedURL, err := h.fileSvc.CreateUploadURL(ctx, req)
	if err != nil {
		common.ToAPIResponse(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	common.ToAPIResponse(c, http.StatusOK, "Generate upload presigned url successfully", gin.H{
		"presigned_url": presignedURL,
	})
}

func (h *FileHandler) ViewFile(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	objectKey := c.Query("key")
	if objectKey == "" {
		common.ToAPIResponse(c, http.StatusBadRequest, common.ErrInvalidQuery.Error(), nil)
		return
	}

	presignedURL, err := h.fileSvc.CreateViewURL(ctx, objectKey)
	if err != nil {
		switch err {
		case common.ErrFileNotFound:
			common.ToAPIResponse(c, http.StatusNotFound, err.Error(), nil)
		default:
			common.ToAPIResponse(c, http.StatusInternalServerError, "internal server error", nil)
		}
		return
	}

	common.ToAPIResponse(c, http.StatusOK, "Generate view presigned url successfully", gin.H{
		"presigned_url": presignedURL,
	})
}
