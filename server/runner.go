package server

import (
	"github.com/Ofla/TODO/config"
	"github.com/Ofla/TODO/errors"
	"github.com/Ofla/TODO/server/services"
	"github.com/sirupsen/logrus"
)

// SvcRunner holds start function te be implemented by a runner
type Runner interface {
	Start() error
}

// Create creates a runner of type defined in config
func Create(conf *config.Config, logger *logrus.Logger) (Runner, error) {
	var srv Runner

	switch conf.Server.Grpc.Type {
	case "grpc":
		srv = services.NewRunner(*conf, logger)
	default:
		return nil, errors.CreateInternalError(errors.InvalidServerTypeError, conf.Server.Grpc.Host)
	}
	return srv, nil
}
