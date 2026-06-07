package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"urlshortener/internal/api"
	"urlshortener/internal/config"
	"urlshortener/internal/repository"
	"urlshortener/internal/service"
)

func main() {
	cfg := config.Load()
	pgURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DB)
	redisURL := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)
	pgRepo, err := repository.NewURLRepo(pgURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pgRepo.Close()

	redisRepo, err := repository.NewRedisRepo(redisURL)
	if err != nil {
		log.Fatal(err)
	}
	defer redisRepo.Close()

	cache := repository.NewCacheStrategy(redisRepo, pgRepo)

	shortener := service.NewShortener(pgRepo, cfg.BaseUrl)
	redirect := service.NewRedirect(cache)

	handlers := api.NewHandlers(shortener, redirect, cfg.BaseUrl)

	router := api.SetupRouter(handlers)

	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: router,
	}

	go func() {
		log.Printf("Listening on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exited")
}
