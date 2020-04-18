package config

import (
	"github.com/jinzhu/configor"
	"path"
	"runtime"
)

// Config is a struct which represents the configuration
type Config struct {
	Server   Server
	Database Database
}
type Server struct {
	Grpc Grpc
	Http Http
}
type Grpc struct {
	Type string
	Host string
	Network string
}

type Http struct {
	Host string
}
type Database struct {
	Type      string
	Tablename string
}

// LoadConfiguration is a func which will load the configuration from config.json file.
func LoadConfiguration() (*Config, error) {
	configFilePath := "../config.yml"
	config := configor.New(&configor.Config{})
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), configFilePath)
	conf := new(Config)
	err := config.Load(conf, filepath)
	return conf, err
}
