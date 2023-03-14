package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liangtengfei/beer-admin/api"
	"github.com/liangtengfei/beer-admin/api/router"
	_ "github.com/liangtengfei/beer-admin/docs"
	"github.com/liangtengfei/beer-admin/internal/conf"
	"github.com/liangtengfei/beer-admin/pkg/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type HTTPServer struct {
	Engine *gin.Engine
}

func NewHTTPServer(
	conf *conf.Server,
	ua *api.SysUserAPI,
) *HTTPServer {
	r := gin.Default()

	// 中间件
	r.Use(middleware.Cors())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 接口文档服务
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.InitSysUserRouter(r, ua)

	return &HTTPServer{
		Engine: r,
	}
}
