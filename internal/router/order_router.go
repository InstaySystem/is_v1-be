package router

import (
	"github.com/InstaySystem/is-be/internal/handler"
	"github.com/InstaySystem/is-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func OrderRouter(rg *gin.RouterGroup, hdl *handler.OrderHandler, authMid *middleware.AuthMiddleware) {
	admin := rg.Group("/admin", authMid.IsAuthentication(), authMid.HasDepartment("reception"))
	{
		admin.POST("/orders/rooms", hdl.CreateOrderRoom)

		admin.GET("/orders/rooms/:id", hdl.GetOrderRoomByID)
	}

	admin = rg.Group("/admin", authMid.IsAuthentication())
	{
		admin.GET("/orders/services", hdl.GetOrderServicesForAdmin)

		admin.GET("/orders/services/:id", hdl.GetOrderServiceByID)
	}

	rg.POST("/orders/rooms/verify", hdl.VerifyOrderRoom)

	guest := rg.Group("/orders", authMid.HasGuestToken())
	{
		guest.POST("/services", hdl.CreateOrderService)

		guest.GET("/services/:code", hdl.GetOrderServiceByCode)

		guest.PUT("/services/:id", hdl.UpdateOrderServiceForGuest)
	}
}