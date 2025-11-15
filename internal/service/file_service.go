package service

import (
	"context"

	"github.com/InstaySystem/is-be/internal/types"
)

type FileService interface {
	CreateUploadURLs(ctx context.Context, req types.UploadPresignedURLsRequest) ([]*types.UploadPresignedURLResponse, error)

	CreateViewURLs(ctx context.Context, req types.ViewPresignedURLsRequest) ([]*types.ViewPresignedURLResponse, error)
}