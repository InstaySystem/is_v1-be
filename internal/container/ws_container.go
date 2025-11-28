package container

import (
	"github.com/InstaySystem/is-be/internal/handler"
	"github.com/InstaySystem/is-be/internal/hub"
)

type WSContainer struct {
	Hdl *handler.WSHandler
}

func NewWSContainer(wsHub *hub.WSHub) *WSContainer {
	hdl := handler.NewWSHandler(wsHub)
	return &WSContainer{hdl}
}