package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alibek-dzhukaev/task-flow/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	db, err := config.InitDatabase(cfg)
	if err != nil {
		log.Fatalf("database error: %v", err)
	}

	redis, err := config.InitRedis(cfg)
	if err != nil {
		log.Fatalf("redis error: %v", err)
	}

	_ = db
	_ = redis

	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: nil,
	}

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		log.Println("shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}()
	log.Printf("server started on port %s", cfg.ServerPort)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
