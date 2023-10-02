package app

import (
	"github.com/rs/zerolog/log"
	"paru.net/gosimpleapp/internal/configuration"
	"paru.net/gosimpleapp/internal/database"
	"paru.net/gosimpleapp/internal/service"
)

func Start() error {
	// getting configuration
	if err := configuration.LoadAppConfig(); err != nil {
		log.Error().Err(err).Msg("no environment config found, using default config")
		return err
	}

	// Setting up dependencies
	db := database.NewMapDb()

	// Setting up services
	fooSvc := service.NewFooService(db)

	// doing something with it...
	fooSvc.DoSomething()

	// TODO: gracefully shutdown
	return nil

}
