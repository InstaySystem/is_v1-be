package repository

import (
	"context"

	"github.com/InstaySystem/is-be/internal/model"
	"github.com/InstaySystem/is-be/internal/types"
	"gorm.io/gorm"
)

type Notification interface {
	CreateNotificationTx(ctx context.Context, tx *gorm.DB, notification *model.Notification) error

	UpdateReadNotificationsByContentIDAndTypeAndReceiver(ctx context.Context, contentID int64, contentType, receiver string, updateData map[string]any) error

	FindAllUnReadNotificationsByContentIDAndTypeAndReceiver(ctx context.Context, staffID, contentID int64, contentType, receiver string) ([]*model.Notification, error)

	CreateNotificationStaffs(ctx context.Context, notificationStaffs []*model.NotificationStaff) error

	FindAllUnReadNotificationsByDepartmentID(ctx context.Context, staffID, departmentID int64) ([]*model.Notification, error)

	CountUnReadNotificationsByDepartmentID(ctx context.Context, userID, departmentID int64) (int64, error)

	FindAllNotificationsByDepartmentIDWithStaffsReadPaginated(ctx context.Context, query types.NotificationPaginationQuery, staffID, departmentID int64) ([]*model.Notification, int64, error)
}