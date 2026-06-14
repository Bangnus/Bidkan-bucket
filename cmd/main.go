package main

import (
	"log"

	"bidkan-bucket/internal/app/initial"
	"bidkan-bucket/internal/config"
)

func main() {
	// 1. Load Config
	cfg := config.LoadConfig()

	// 2. Initialize App
	app := initial.InitializeApp(cfg)

	// 3. Start Server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := app.Listen(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
