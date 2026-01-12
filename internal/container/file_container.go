package container

import (
	"github.com/InstaySystem/is_v1-be/internal/config"
	"github.com/InstaySystem/is_v1-be/internal/handler"
	"github.com/InstaySystem/is_v1-be/internal/initialization"
	svcImpl "github.com/InstaySystem/is_v1-be/internal/service/implement"
	"go.uber.org/zap"
)

type FileContainer struct {
	Hdl *handler.FileHandler
}

func NewFileContainer(
	cfg *config.Config,
	s3 *initialization.S3,
	logger *zap.Logger,
) *FileContainer {
	svc := svcImpl.NewFileService(s3.Client, s3.Presigner, cfg, logger)
	hdl := handler.NewFileHandler(svc)

	return &FileContainer{hdl}
}
