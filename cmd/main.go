package main

import (
	"os"

	"github.com/jwo_auth/internal/api"
	"github.com/jwo_auth/internal/common"
)

type Application struct {
	apiServer *api.ApiServer
	//TODO: Add additional dependencies here.
}

func (app Application) Start() {
	app.apiServer.Start()
}

func NewApplication(config *common.Config) *Application {
	return &Application{
		apiServer: api.NewApiServer(config),
	}
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
