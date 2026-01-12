package container

import (
	"github.com/InstaySystem/is_v1-be/internal/handler"
	"github.com/InstaySystem/is_v1-be/internal/repository"
	svcImpl "github.com/InstaySystem/is_v1-be/internal/service/implement"
	"github.com/InstaySystem/is_v1-be/pkg/snowflake"
	"go.uber.org/zap"
)

type RoomContainer struct {
	Hdl *handler.RoomHandler
}

func NewRoomContainer(
	roomRepo repository.RoomRepository,
	sfGen snowflake.Generator,
	logger *zap.Logger,
) *RoomContainer {
	svc := svcImpl.NewRoomService(roomRepo, sfGen, logger)
	hdl := handler.NewRoomHandler(svc)

	return &RoomContainer{hdl}
}
