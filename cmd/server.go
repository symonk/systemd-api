package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
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

func newServer(lc fx.Lifecycle, cfg *config.Config) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	engine := gin.New()

	// Set up logging middleware?

	// Set up metrics?

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.)
	}

}
