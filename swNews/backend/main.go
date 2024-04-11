package main

import (
	"backend/bootstrap"
	"backend/pkg/conf"
	"backend/pkg/logger"
	"backend/routes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 项目初始化
	bootstrap.Init()
	api := routes.Init()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.ServerConf.Port),
		Handler: api,
	}
	go func() {
		fmt.Printf("Server started at :%d%s \n", conf.ServerConf.Port, conf.ServerConf.Prefix)
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("server error", err)
			return
		}
	}()

	// 平滑关机
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("server shutdown error", err)
	}
	err := bootstrap.Shutdown()
	if err != nil {
		logger.Error("cache shutdown error", err)
	}
	logger.Info("Server exiting")
}
