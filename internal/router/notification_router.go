package router

import (
	"github.com/InstaySystem/is-be/internal/handler"
	"github.com/InstaySystem/is-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NotificationRouter(rg *gin.RouterGroup, hdl *handler.NotificationHandler, authMid *middleware.AuthMiddleware) {
	admin := rg.Group("/admin", authMid.IsAuthentication())
	{
		admin.GET("/notifications", hdl.GetNotificationsForAdmin)

		admin.GET("/notifications/unread-count", hdl.CountUnreadNotificationsForAdmin)
	}
}