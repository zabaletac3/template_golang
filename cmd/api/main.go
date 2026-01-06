package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/eren_dev/go_server/internal/app"
	"github.com/eren_dev/go_server/internal/config"
	"github.com/eren_dev/go_server/internal/platform/logger"
)

func main() {
	cfg := config.Load()

	log := logger.NewSlogLogger(cfg.Env)
	
	logger.SetDefault(log)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	server := app.NewServer(cfg)

	logger.Default().Info(context.Background(),
	"server_running",
	"port", cfg.Port,
	"env", cfg.Env,
)

	go func (){
		if err := server.Start(); err != nil {
			logger.Default().Error(context.Background(), "server_error", "error", err)
		}
	}()

	<-ctx.Done()

	server.Shutdown(context.Background())

	logger.Default().Info(context.Background(), "server_stopped")

	os.Exit(0)
}