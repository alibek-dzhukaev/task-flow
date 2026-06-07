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
	"github.com/alibek-dzhukaev/task-flow/internal/handler"
	"github.com/alibek-dzhukaev/task-flow/internal/model"
	"github.com/alibek-dzhukaev/task-flow/internal/repository"
	"github.com/alibek-dzhukaev/task-flow/internal/router"
	"github.com/alibek-dzhukaev/task-flow/internal/service"
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

	redisClient, err := config.InitRedis(cfg)
	if err != nil {
		log.Fatalf("redis error: %v", err)
	}

	_ = redisClient

	if err := db.AutoMigrate(
		&model.User{},
		&model.Project{},
		&model.Task{},
		&model.Label{},
		&model.Comment{},
		&model.Activity{},
	); err != nil {
		log.Fatalf("database migration error: %v", err)
	}

	tokenTTl, _ := time.ParseDuration(cfg.JWTAccessTokenExpiry)
	userRepo := repository.NewUserRepository(db)
	authSvc := service.NewAuthService(userRepo, cfg.JWTSecret, tokenTTl)
	authHandler := handler.NewAuthHandler(authSvc)

	r := router.New(authHandler, cfg.JWTSecret)

	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: r,
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
