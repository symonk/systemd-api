package main

import (
	"log"
	"os"

	"github.com/symonk/systemd-api/cmd/server"
)

func main() {
	if err := server.Execute(); err != nil {
		log.Printf("Failed launching the root command. %v", err)
		os.Exit(1)
	}
}
