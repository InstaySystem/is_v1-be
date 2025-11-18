package service

import (
	"context"

	"github.com/InstaySystem/is-be/internal/model"
	"github.com/InstaySystem/is-be/internal/types"
)

type RoomService interface {
	CreateRoomType(ctx context.Context, userID int64, req types.CreateRoomTypeRequest) error

	GetRoomTypesForAdmin(ctx context.Context) ([]*model.RoomType, error)

	UpdateRoomType(ctx context.Context, roomTypeID, userID int64, req types.UpdateRoomTypeRequest) error

	DeleteRoomType(ctx context.Context, roomTypeID int64) error

	CreateRoom(ctx context.Context, userID int64, req types.CreateRoomRequest) error
}