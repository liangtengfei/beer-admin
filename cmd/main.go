package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/liangtengfei/beer-admin/internal/conf"
	"github.com/liangtengfei/beer-admin/internal/server"
	"github.com/liangtengfei/beer-admin/pkg/logx"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "beer.admin.service"
	// Version is the version of the compiled software.
	Version string = "v1.0.0"
	// flagconf is the config flag.
	flagconf string

	DevMode = true
)

type App struct {
	Name   string
	engine *gin.Engine
}

func newApp(httpServer *server.HTTPServer) *App {
	return &App{
		Name:   Name,
		engine: httpServer.Engine,
	}
}

func init() {
	flag.StringVar(&flagconf, "conf", "../configs/", "config path, eg: -conf config.yaml")
	flag.BoolVar(&DevMode, "mode", true, "development mode, eg: -mode true")
}

//	@title			BeerAdmin巡防管理系统
//	@version		1.0.0
//	@description	BeerAdmin巡防管理系统-基础管理服务
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8000
//	@BasePath	/api/sys

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @tokenUrl					http://localhost:8000/api/sys/login
// @description				OAuth protects our entity endpoints
func main() {
	flag.Parse()
	logger := logx.NewLogger(DevMode, Name)
	defer logger.Sync()
	sugar := logger.Sugar().Named("BEER-ADMIN")

	// 读取配置文件信息
	config := conf.NewConfig(flagconf)

	config.Server.Casbin.ModelPath = flagconf + config.Server.Casbin.ModelPath

	// 应用初始化
	app, cleanup, err := initApp(config.Server, config.Registry, config.Data, config.App, sugar)
	if err != nil {
		log.Fatalf("初始化应用失败：%v", err)
	}
	defer cleanup()

	// 优雅地重启或停止
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    config.Server.Http.Addr,
		Handler: app.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("启动服务失败: %v\n", err)
		}
	}()

	log.Printf("[HTTP] 服务启动: %s 服务名称：%s 服务版本：%s\n", srv.Addr, Name, Version)

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
	// 停止
}
