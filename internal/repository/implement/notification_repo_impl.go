package implement

import (
	"context"

	"github.com/InstaySystem/is-be/internal/model"
	"github.com/InstaySystem/is-be/internal/repository"
	"gorm.io/gorm"
)

type notificationRepoImpl struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) repository.Notification {
	return &notificationRepoImpl{db}
}

func (r *notificationRepoImpl) CreateNotificationTx(ctx context.Context, tx *gorm.DB, notification *model.Notification) error {
	return tx.WithContext(ctx).Create(notification).Error
}

func (r *notificationRepoImpl) CreateNotificationStaffs(ctx context.Context, notificationStaffs []*model.NotificationStaff) error {
	return r.db.WithContext(ctx).Create(notificationStaffs).Error
}

func (r *notificationRepoImpl) UpdateReadNotificationsByContentIDAndTypeAndReceiver(ctx context.Context, contentID int64, contentType, receiver string, updateData map[string]any) error {
	return r.db.WithContext(ctx).Model(&model.Notification{}).Where("content_id = ? AND type = ? AND receiver = ? AND is_read = false", contentID, contentType, receiver).Updates(updateData).Error
}

func (r *notificationRepoImpl) FindAllUnReadNotificationsByContentIDAndTypeAndReceiver(ctx context.Context, staffID, contentID int64, contentType, receiver string) ([]*model.Notification, error) {
	var notifications []*model.Notification
	if err := r.db.WithContext(ctx).Preload("StaffRead", "staff_id = ?").Where("content_id = ? AND type = ? AND receiver = ?", contentID, contentType, receiver).Where("id NOT IN (?)",
		r.db.Model(&model.NotificationStaff{}).
			Select("notification_id").
			Where("staff_id = ?", staffID),
	).Order("created_at DESC").Find(&notifications).Error; err != nil {
		return nil, err
	}

	return notifications, nil
}
