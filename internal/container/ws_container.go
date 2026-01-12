package container

import (
	"github.com/InstaySystem/is_v1-be/internal/handler"
	"github.com/InstaySystem/is_v1-be/internal/hub"
)

type WSContainer struct {
	Hdl *handler.WSHandler
}

func NewWSContainer(wsHub *hub.WSHub) *WSContainer {
	hdl := handler.NewWSHandler(wsHub)
	return &WSContainer{hdl}
}
