package webserver

import (
	"context"
	"net/http"

	"github.com/lyonnee/go-gin-template/api"
	"github.com/lyonnee/go-gin-template/infra/config"
	"github.com/lyonnee/go-gin-template/infra/log"
	"go.uber.org/zap"
)

var srv *http.Server

func Run(env string) {
	log.Debug("Web API Service run")

	// 获取路由器实例
	r := newRouter(env)
	// 初始化handler(注册路由)
	api.Register(r)

	srv = &http.Server{
		Addr:    config.App().Port,
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Web API Service Not Run", zap.Error(err))
	}
}

func Shutdown(ctx context.Context) {
	srv.Shutdown(ctx)
}
