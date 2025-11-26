package repository

import (
	"context"

	"github.com/InstaySystem/is-be/internal/model"
	"gorm.io/gorm"
)

type Notification interface {
	CreateNotificationTx(ctx context.Context, tx *gorm.DB, notification *model.Notification) error

	UpdateReadNotificationsByContentIDAndTypeAndReceiver(ctx context.Context, contentID int64, contentType, receiver string, updateData map[string]any) error

	FindAllUnReadNotificationsByContentIDAndTypeAndReceiver(ctx context.Context, staffID, contentID int64, contentType, receiver string) ([]*model.Notification, error)

	CreateNotificationStaffs(ctx context.Context, notificationStaffs []*model.NotificationStaff) error
}