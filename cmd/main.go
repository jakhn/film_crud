package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jakhn/film_crud/api"
	"github.com/jakhn/film_crud/config"
	"github.com/jakhn/film_crud/storage/postgres"
)

func main() {

	cfg := config.Load()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	storage, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer storage.CloseDB()

	api.SetUpApi(&cfg,r, storage)

	log.Printf("Listening port %v...\n", cfg.HTTPPort)
	err = r.Run(cfg.HTTPPort)
	if err != nil {
		panic(err)
	}
}
