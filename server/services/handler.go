package services

import (
	"fmt"
	"github.com/Ofla/TODO/config"
	"github.com/Ofla/TODO/database"
	svc "github.com/Ofla/TODO/proto/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type SvcRunner struct {
	database.DbHandler
	config.Config
	*logrus.Logger
}

// NewRunner create a runner by creating its all dependencies (database, s3, sqs)
func NewRunner(conf config.Config, logger *logrus.Logger) *SvcRunner {
	dbHandler, err := database.CreateDbHandler(&conf.Database, logger)
	if err != nil {
		logger.Panicf("Fail to create db handler :%v", err)
	}
	return &SvcRunner{
		DbHandler:      dbHandler,
		Config:         conf,
		Logger:         logger,
	}

}

// Start starts the runner , this method will be called in main function to run the server
func (srv *SvcRunner) Start() error {
	srv.Logger.Println("Starting server...")
	// create listener
	listener, err := net.Listen(
		"tcp",
		fmt.Sprintf(":%v", srv.Config.Server.Grpc.Network),
	)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	svc.RegisterTodoServiceServer(server, srv)

	return server.Serve(listener)
}