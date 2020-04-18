package services

import (
	"github.com/Ofla/TODO/config"
	hlp "github.com/Ofla/TODO/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ttRunner = []struct {
	Name string
	conf config.Config
}{
	{"unsupported Database", config.Config{
		Database: config.Database{
			Type:      "wrongType",
			Tablename: "todo",
		},
	}},
}

func TestNewRunner(t *testing.T) {
	for _, testCase := range ttRunner {
		log := hlp.GetLogger()
		t.Run(testCase.Name, func(t *testing.T) {
			assert.Panics(t, func() { _ = NewRunner(testCase.conf, log) }, "Creating runner should panic")
		})
	}
}
