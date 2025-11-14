package implement

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/InstaySystem/is-be/internal/common"
	"github.com/InstaySystem/is-be/internal/config"
	"github.com/InstaySystem/is-be/internal/service"
	"github.com/InstaySystem/is-be/internal/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type fileSvcImpl struct {
	client    *s3.Client
	presigner *s3.PresignClient
	cfg       *config.Config
	logger    *zap.Logger
}

func NewFileService(client *s3.Client, presigner *s3.PresignClient, cfg *config.Config, logger *zap.Logger) service.FileService {
	return &fileSvcImpl{
		client,
		presigner,
		cfg,
		logger,
	}
}

func (s *fileSvcImpl) CreateUploadURL(ctx context.Context, req types.PresignedURLRequest) (string, error) {
	objectKey := fmt.Sprintf("%s/%s-%s", s.cfg.S3.Folder, uuid.NewString(), common.GenerateSlug(req.FileName))

	presignedRes, err := s.presigner.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.cfg.S3.Bucket),
		Key:         aws.String(objectKey),
		ContentType: aws.String(req.ContentType),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = 15 * time.Minute
	})
	if err != nil {
		s.logger.Error("generate upload presigned URL failed", zap.Error(err))
		return "", err
	}

	return presignedRes.URL, nil
}

func (s *fileSvcImpl) CreateViewURL(ctx context.Context, objectKey string) (string, error) {
	if _, err := s.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.cfg.S3.Bucket),
		Key:    aws.String(objectKey),
	}); err != nil {
		var keyNotFound *s3Types.NotFound
		if errors.As(err, &keyNotFound) {
			return "", common.ErrFileNotFound
		}
		s.logger.Error("file check failed", zap.Error(err))
		return "", err
	}

	presignedReq, err := s.presigner.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.cfg.S3.Bucket),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = 15 * time.Minute
	})
	if err != nil {
		s.logger.Error("generate view presigned URL failed", zap.Error(err))
		return "", err
	}

	return presignedReq.URL, nil
}
