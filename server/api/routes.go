package api

import (
	"server/api/world"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	worldHandler *world.WorldHandler,
) {
	api := router.Group("/api/v1")

	worlds := api.Group("/worlds")
	{
		worlds.POST("", worldHandler.CreateWorld)
	}

}
