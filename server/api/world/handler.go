package world

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WorldHandler struct {
	service WorldService
}

func (h *WorldHandler) CreateWorld(c *gin.Context) {
	var req CreateWorldRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	world, err := h.service.CreateWorld(c.Request.Context(), req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	WorldResponse := WorldResponse{
		Grid:   world,
		Height: len(world),
		Width:  len(world[0]),
	}

	c.JSON(http.StatusOK, WorldResponse)
}

func NewHandler(service WorldService) *WorldHandler {
	return &WorldHandler{
		service: service,
	}
}
