package main

// import "github.com/gin-gonic/gin"
import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"rt-msg-carrier/configs"
	"rt-msg-carrier/log"
	"rt-msg-carrier/router"
)

func main() {
	r := router.SetupRouter()
	logger := log.NewLogger()
	srv_cfg := configs.Get().Server
	server := http.Server{
		Addr:         srv_cfg.Addr,
		Handler:      r,
		ReadTimeout:  time.Duration(srv_cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(srv_cfg.WriteTimeout) * time.Second,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("server listen err: %s", err)
		}
	}()
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("server shutdown error")
	}
	logger.Println("server exiting...")

}
