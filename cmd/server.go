package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"honnef.co/go/tools/config"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		runApplication()
	},
}

func runApplication() {
	config, err := config.Load(configFile)
	if err != nil {
		log.Fatal(err)
	}

}

func newServer(lifecycle fx.Lifecycle, cfg *config.Config) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	engine := gin.New()

	// Set up logging middleware?

	// Set up metrics?

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.ServerConfig.Port),
		Handler:      engine,
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
		},
		OnStop: func(ctx context.Context) error {
			logging.FromContext(ctx).Info("Stopping rest api server")
			return server.Shutdown(ctx)
		},
	})
	return engine
}
