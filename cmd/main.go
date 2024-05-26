package main

import (
	"chatbotmrrobotdumbazz/delivery"
	"chatbotmrrobotdumbazz/goconfig"
	"chatbotmrrobotdumbazz/logger"
	"chatbotmrrobotdumbazz/server"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/exp/slog"
)

func main() {
	cfg := goconfig.MustLoad()
	log := logger.SetupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))
	handlers := delivery.NewHandler(log)
	log.Info("initializing server", slog.String("address", cfg.Address))
	log.Debug("logger debug mode enabled")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	server := new(server.Server)
	go func() {
		if err := server.Start(":7070", handlers.Handlers()); err != nil {
			log.With(err)
			return
		}
		log.Debug("server started")
	}()
	if err := server.Shutdown(ctx); err != nil {
		log.With(err.Error())
		return
	}
}
