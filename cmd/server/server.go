package server

import (
	"github.com/Ofla/TODO/config"
	"github.com/Ofla/TODO/errors"
	hlp "github.com/Ofla/TODO/helpers"
	"github.com/Ofla/TODO/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := hlp.GetLogger()
	logger.Debug("Starting provider service...")
	// get the configuration
	conf, err := config.LoadConfiguration()
	if err != nil {
		logger.Fatalf(errors.CreateInternalError(err).Error())
	}
	logger.Debug("Configuration parsed successfully")

	srv := getServer(conf, logger)

	err = srv.Start()
	if err != nil {
		logger.Fatal(errors.CreateInternalError(err).Error())
	}
	logger.Debug("Exiting...")
}

func getServer(conf *config.Config, logger *logrus.Logger) server.Runner {
	srv, err := server.Create(conf, logger)
	if err != nil {
		logger.Fatal(errors.CreateInternalError(err).Error())
	}
	logger.Debug("Accepting requests...")
	return srv
}
