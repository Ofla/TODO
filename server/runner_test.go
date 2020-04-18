package server_test

import (
	"github.com/Ofla/TODO/config"
	hlp "github.com/Ofla/TODO/helpers"
	"github.com/Ofla/TODO/server"
	"github.com/sirupsen/logrus"
	"testing"
)

var (
	conf *config.Config
	log  *logrus.Logger
)

func init() {
	conf, _ = config.LoadConfiguration()
	log = hlp.GetLogger()
}

func TestCreate(t *testing.T) {
	t.Run("failure in runner creation", func(t *testing.T) {
		serverRunner := config.Server{
			Grpc: config.Grpc{
				Type: "wrongType",
			},
		}
		wrongConf := config.Config{
			Server: serverRunner,
		}
		runner, err := server.Create(&wrongConf, log)
		if err == nil || runner != nil {
			t.Error("expected error")
		}
	})
	t.Run("success runner creation", func(t *testing.T) {
		_, err := server.Create(conf, log)
		if err != nil {
			t.Errorf("expected error, got %v", err)
		}
	})
}
