package container

import (
	"github.com/InstaySystem/is-be/internal/config"
	"github.com/InstaySystem/is-be/internal/initialization"
	"go.uber.org/zap"
)

type Container struct {
	FileContainer *FileContainer
}

func NewContainer(cfg *config.Config, s3 *initialization.S3, logger *zap.Logger) *Container {
	fileCtn := NewFileContainer(cfg, s3, logger)
	return &Container{fileCtn}
}