package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/lyonnee/go-gin-template/infra/config"
	"github.com/lyonnee/go-gin-template/infra/database"
	"github.com/lyonnee/go-gin-template/infra/log"
	"github.com/lyonnee/go-gin-template/webserver"
	"github.com/spf13/cobra"
)

var (
	flagEnv string

	rootCmd = &cobra.Command{
		Use:   "go-gin-template",
		Short: "go-gin-template is a restful api program",
		Run: func(cmd *cobra.Command, args []string) {
			entry()
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&flagEnv, "environment", "e", "dev", "default is dev,select dev/test/prod")
}

// 程序真正执行入口
func entry() {
	// 初始化配置
	config.Load(flagEnv)
	// 初始化日志组件
	log.Initialize()

	// 连接数据库
	database.Connect()

	// 启动web api服务
	webserver.Run(flagEnv)

	log.Info("Server Running ...")

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Info("Server Shutdown ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	webserver.Shutdown(ctx)
	database.Disconnect()
	log.Sync()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
