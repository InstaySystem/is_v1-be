package container

import (
	"github.com/InstaySystem/is_v1-be/internal/handler"
	"github.com/InstaySystem/is_v1-be/internal/repository"
	svcImpl "github.com/InstaySystem/is_v1-be/internal/service/implement"
	"github.com/InstaySystem/is_v1-be/pkg/snowflake"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type NotificationContainer struct {
	Hdl *handler.NotificationHandler
}

func NewNotificationContainer(
	db *gorm.DB,
	notificationRepo repository.Notification,
	logger *zap.Logger,
	sfGen snowflake.Generator,
) *NotificationContainer {
	svc := svcImpl.NewNotificationService(db, notificationRepo, logger, sfGen)
	hdl := handler.NewNotificationHandler(svc)

	return &NotificationContainer{hdl}
}
