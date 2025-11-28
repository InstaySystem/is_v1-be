package router

import (
	"github.com/InstaySystem/is-be/internal/handler"
	"github.com/InstaySystem/is-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func WSRouter(rg *gin.RouterGroup, hdl *handler.WSHandler, authMid *middleware.AuthMiddleware) {
	rg.GET("/ws", authMid.IsClient(), hdl.ServeWS)
}