package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/symonk/systemd-api/pkg/logging"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/symonk/systemd-api/internal/config"
	"go.uber.org/fx"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		runApplication()
	},
}

func runApplication() {
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}
	logging.SetConfig(&logging.Config{
		Encoding: cfg.LoggingConfig.Encoding,
		// Todo: Fix level serialization.
		Level:       10,
		Development: cfg.LoggingConfig.Development,
	})
	defer logging.DefaultLogger().Sync()

	app := fx.New(
		fx.Supply(cfg),
		fx.Supply(logging.DefaultLogger().Desugar()),
		fx.Provide(newServer, newEchohandler),
		fx.Invoke(func(*http.Server) {}),
	)
	app.Run()

}

func newServer(lifecycle fx.Lifecycle, cfg *config.Config) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.New()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.ServerConfig.Port),
		Handler:      router,
		ReadTimeout:  cfg.ServerConfig.ReadTimeout,
		WriteTimeout: cfg.ServerConfig.WriteTimeout,
	}
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logging.FromContext(ctx).Infof("Starting rest api server :%d", cfg.ServerConfig.Port)
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logging.DefaultLogger().Errorw("failed to close http server", "err", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logging.FromContext(ctx).Info("Stopping rest api server")
			return server.Shutdown(ctx)
		},
	})
	return router
}

type EchoHandler struct{}

func newEchohandler() *EchoHandler {
	return &EchoHandler{}
}

func (*EchoHandler) serveHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		fmt.Fprintln(os.Stderr, "failed to handle request", err)
	}
}
