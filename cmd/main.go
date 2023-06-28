package main

import (
	"time"

	"github.com/jwo_auth/internal/api"
	"github.com/jwo_auth/internal/database"
)

type Config struct {
	serverEndpoint                       string
	serverReadTimeout                    time.Duration
	serverWriteTimeout                   time.Duration
	serverHeaderAccessControlAllowOrigin string
	jwtSigningKey                        string
	jwtValidityInHours                   time.Duration
}

type Application struct {
	server   api.Server
	database database.Database
}

func (app Application) Start() {}

func NewApplication(config *Config) *Application {
	return &Application{}
}

func main() {
	config := &Config{}
	app := NewApplication(config)
	app.Start()
}
