package main

import (
	"server/api"
	"server/api/world"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := api.LoadConfig()

	worldService := world.NewService()
	worldHandler := world.NewHandler(worldService)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	api.RegisterRoutes(router, worldHandler)

	router.Run(cfg.HTTPAddress)

}
