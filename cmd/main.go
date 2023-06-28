package main

import (
	"os"

	"github.com/jwo_auth/internal/api"
	"github.com/jwo_auth/internal/common"
	"github.com/jwo_auth/internal/database"
)

type Application struct {
	server   api.Server
	database database.Database
}

func (app Application) Start() {}

func NewApplication(config *common.Config) *Application {
	return &Application{}
}

func main() {
	config, err := common.ReadConfig()

	if err != nil {
		os.Exit(1)
	} else {
		app := NewApplication(config)
		app.Start()
	}
}
