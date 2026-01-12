package container

import (
	"github.com/InstaySystem/is_v1-be/internal/handler"
	"github.com/InstaySystem/is_v1-be/internal/provider/mq"
	"github.com/InstaySystem/is_v1-be/internal/repository"
	svcImpl "github.com/InstaySystem/is_v1-be/internal/service/implement"
	"github.com/InstaySystem/is_v1-be/pkg/snowflake"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ServiceContainer struct {
	Hdl *handler.ServiceHandler
}

func NewServiceContainer(
	db *gorm.DB,
	serviceRepo repository.ServiceRepository,
	sfGen snowflake.Generator,
	logger *zap.Logger,
	mqProvider mq.MessageQueueProvider,
) *ServiceContainer {
	svc := svcImpl.NewServiceService(serviceRepo, db, sfGen, logger, mqProvider)
	hdl := handler.NewServiceHandler(svc)

	return &ServiceContainer{hdl}
}
