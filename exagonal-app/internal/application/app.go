package application

import (
	"log"

	"paru.net/exagonalapp/internal/adapters/db"
	"paru.net/exagonalapp/internal/application/api"
	"paru.net/exagonalapp/internal/application/domain"
	"paru.net/exagonalapp/internal/config"
)

func Start() {
	// getting configuration
	if err := config.LoadAppConfig(); err != nil {
		log.Fatalf("no environment config found, using default config")
	}
	dbAdapter, err := db.NewAdapter(config.Application.DataSource.Url)
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	api := api.NewApplication(dbAdapter)
	orderItems := []domain.OrderItem{
		{
			ProductCode: "daje",
			UnitPrice:   1.0,
			Quantity:    1,
		},
	}
	or1 := domain.NewOrder(1, orderItems)
	api.PlaceOrder(or1)

}
