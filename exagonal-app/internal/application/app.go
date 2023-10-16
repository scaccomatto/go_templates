package application

import (
	"log"

	"paru.net/exagonalapp/internal/adapters/db"
	"paru.net/exagonalapp/internal/application/api"
	"paru.net/exagonalapp/internal/config"
)

func Start() {
	// getting configuration
	if err := config.LoadAppConfig(); err != nil {
		log.Fatalf("no environment config found, using default config")
	}
	dbAdapter, err := db.NewAdapter(config.Application.DataSource.HostPort)
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	api.NewApplication(dbAdapter)

}
