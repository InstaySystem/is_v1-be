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

type RequestContainer struct {
	Hdl *handler.RequestHandler
}

func NewRequestContainer(
	db *gorm.DB,
	requestRepo repository.RequestRepository,
	orderRepo repository.OrderRepository,
	notificationRepo repository.Notification,
	sfGen snowflake.Generator,
	logger *zap.Logger,
	mqProvider mq.MessageQueueProvider,
) *RequestContainer {
	svc := svcImpl.NewRequestService(db, requestRepo, orderRepo, notificationRepo, sfGen, logger, mqProvider)
	hdl := handler.NewRequestHandler(svc)

	return &RequestContainer{hdl}
}
