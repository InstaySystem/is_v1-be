package container

import (
	"github.com/InstaySystem/is-be/internal/handler"
	"github.com/InstaySystem/is-be/internal/repository"
	svcImpl "github.com/InstaySystem/is-be/internal/service/implement"
	"github.com/InstaySystem/is-be/pkg/snowflake"
	"go.uber.org/zap"
)

type NotificationContainer struct {
	Hdl *handler.NotificationHandler
}

func NewNotificationContainer(
	notificationRepo repository.Notification,
	logger *zap.Logger,
	sfGen snowflake.Generator,
) *NotificationContainer {
	svc := svcImpl.NewNotificationService(notificationRepo, logger, sfGen)
	hdl := handler.NewNotificationHandler(svc)

	return &NotificationContainer{hdl}
}
