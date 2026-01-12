package router

import (
	"github.com/InstaySystem/is_v1-be/internal/handler"
	"github.com/InstaySystem/is_v1-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func BookingRouter(rg *gin.RouterGroup, hdl *handler.BookingHandler, authMid *middleware.AuthMiddleware) {
	admin := rg.Group("/admin", authMid.IsAuthentication(), authMid.HasDepartment("reception"))
	{
		admin.GET("/bookings", hdl.GetBookings)

		admin.GET("/bookings/:id", hdl.GetBookingByID)

		admin.GET("/sources", hdl.GetSources)
	}
}
