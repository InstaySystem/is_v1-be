package service

import (
	"context"

	"github.com/InstaySystem/is-be/internal/types"
)

type FileService interface {
	CreateUploadURL(ctx context.Context, req types.PresignedURLRequest) (string, error)
}