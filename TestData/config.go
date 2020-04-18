package TestData

import "github.com/Ofla/TODO/config"

// TTCreateHandler represents table test structure of CreateHandler test
type TTCreateHandler struct {
	Name     string
	TodoDb   config.Database
	HasError bool
}

// CreateTTHandler creates table test for CreateHandler test
func CreateTTHandler() []TTCreateHandler {
	return []TTCreateHandler{
		{
			Name: "valid config ",
			TodoDb: config.Database{
				Type:      "memdb",
				Tablename: "todo",
			},
			HasError: false,
		},
		{
			Name: "unsupported db type",
			TodoDb: config.Database{
				Type:      "wrongType",
				Tablename: "todo",
			},
			HasError: true,
		},
	}
}
