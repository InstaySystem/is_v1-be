package container

import (
	"github.com/InstaySystem/is_v1-be/internal/handler"
	"github.com/InstaySystem/is_v1-be/internal/hub"
)

type SSEContainer struct {
	Hdl *handler.SSEHandler
}

func NewSSEContainer(sseHub *hub.SSEHub) *SSEContainer {
	hdl := handler.NewSSEHandler(sseHub)
	return &SSEContainer{hdl}
}
