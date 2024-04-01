package services

import (
	"context"
	"net/http"

	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/gin-gonic/gin"
	"github.com/symonk/systemd-api/internal/config"
	"github.com/symonk/systemd-api/internal/middleware"
	"github.com/symonk/systemd-api/internal/systemd"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) list(c *gin.Context) {
	dbus := c.MustGet("dbus").(*dbus.Conn)
	units, err := dbus.ListUnitsContext(context.Background())
	_ = units
	_ = err
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func RouteV1(cfg *config.Config, h *Handler, r *gin.Engine, dbus *systemd.ServiceInspector) {
	v1 := r.Group("v1/api")
	serviceV1 := v1.Group("services")
	{
		serviceV1.GET("/", h.list)
	}
	serviceV1.Use(middleware.InspectorMiddleware(dbus))
}
