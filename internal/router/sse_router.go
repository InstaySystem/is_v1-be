package router

import (
	"github.com/InstaySystem/is_v1-be/internal/handler"
	"github.com/InstaySystem/is_v1-be/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SSERouter(rg *gin.RouterGroup, hdl *handler.SSEHandler, authMid *middleware.AuthMiddleware) {
	rg.GET("/sse", authMid.IsGuestOrStaffHasDepartment(nil), hdl.ServeSSE)
}
