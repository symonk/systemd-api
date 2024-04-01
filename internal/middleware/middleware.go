package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/symonk/systemd-api/internal/systemd"
)

// InspectorMiddleware registers a new system dbus service
// to the context for interacting via dbus with systemd.
func InspectorMiddleware(dbus *systemd.ServiceInspector) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbus", dbus)
	}
}
