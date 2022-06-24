package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/LyonNee/app-layout/database"
	"github.com/LyonNee/app-layout/ginserver"
	"github.com/LyonNee/app-layout/pkg/config"
	"github.com/LyonNee/app-layout/pkg/log"
	"github.com/spf13/cobra"
)

var (
	flagEnv string

	rootCmd = &cobra.Command{
		Use:   "app-layout",
		Short: "app-layout is a restful api program",
		Run: func(cmd *cobra.Command, args []string) {
			entry()
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&flagEnv, "environment", "e", "dev", "default is dev,select dev or prod")
}

// 程序真正执行入口
func entry() {
	// 初始化配置
	config.Initialize(flagEnv)
	// 初始化日志组件
	log.Initialize()

	// 连接数据库
	database.Connect()

	// 启动服务
	go ginserver.Run()

	log.ZapLogger().Info("Server Running ...")

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.ZapLogger().Info("Server Shutdown ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ginserver.Shutdown(ctx)
	database.Disconnect()
	log.Sync()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
