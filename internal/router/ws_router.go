package router

import (
	"github.com/InstaySystem/is_v1-be/internal/handler"
	"github.com/InstaySystem/is_v1-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func WSRouter(rg *gin.RouterGroup, hdl *handler.WSHandler, authMid *middleware.AuthMiddleware) {
	allowDept := "customer-care"
	rg.GET("/ws", authMid.IsGuestOrStaffHasDepartment(&allowDept), hdl.ServeWS)
}
