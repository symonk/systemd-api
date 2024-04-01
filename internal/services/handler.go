package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/symonk/systemd-api/internal/config"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) services(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func RouteV1(cfg *config.Config, h *Handler, r *gin.Engine) {
	v1 := r.Group("v1/api")
	// Consider route middleware.
	serviceV1 := v1.Group("services")
	serviceV1.Use()
	{
		serviceV1.GET("/services", h.services)
	}
}
