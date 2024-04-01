package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/symonk/systemd-api/pkg/logging"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/symonk/systemd-api/internal/config"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
		Level:       zapcore.Level(cfg.LoggingConfig.Level),
		Development: cfg.LoggingConfig.Development,
	})
	defer logging.DefaultLogger().Sync()

	app := fx.New(
		fx.Supply(cfg),
		fx.Supply(logging.DefaultLogger().Desugar()),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logging.DefaultLogger().Desugar()}
		}),
		fx.StopTimeout(cfg.ServerConfig.GracefulShutdown*time.Second),
		fx.Provide(newServer),
		fx.Invoke(printConfigInfo),
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

func printConfigInfo(cfg *config.Config) {
	b, _ := json.MarshalIndent(&cfg, "", " ")
	logging.DefaultLogger().Infof("config information\n%s", string(b))
}
