package systemd

import (
	"context"

	"github.com/coreos/go-systemd/v22/dbus"
	"go.uber.org/zap"
)

type ServiceInspector struct {
	sysd *dbus.Conn
}

func NewServiceInspector(l *zap.Logger) *ServiceInspector {
	ctx := context.Background()
	dbus, err := dbus.NewSystemdConnectionContext(ctx)
	if err != nil {
		l.Sugar().Fatalf("Cannot establish a connection to systemd over dbus. %w", err)
	}
	return &ServiceInspector{
		sysd: dbus,
	}
}
