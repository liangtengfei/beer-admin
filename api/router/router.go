package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liangtengfei/beer-admin/api"
	"go.uber.org/zap"
)

func InitSysUserRouter(router *gin.Engine, a *api.SysUserAPI) {
	a.Logger = a.Logger.With(zap.String("mobule", "api/user"))
	r := router.Group("/api/sys/user")
	{
		r.GET("/:id", a.GetSysUser)
		r.POST("/", a.AddSysUser)
		r.PUT("/", a.EditSysUser)
		r.DELETE("/:id", a.DelSysUser)
		r.POST("/page", a.PageSysUser)
	}
}
