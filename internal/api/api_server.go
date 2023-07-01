package api

import (
	"net/http"

	"github.com/jwo_auth/internal/common"
	"github.com/jwo_auth/internal/common/logger"
)

type ApiServer struct {
	server *http.Server
}

func NewApiServer(config *common.Config) *ApiServer {
	return &ApiServer{
		server: &http.Server{
			Addr: config.ServerEndpoint,
			//TODO: Add own custom handler here.
			ReadTimeout:  config.ServerReadTimeout,
			WriteTimeout: config.ServerWriteTimeout,
		},
	}
}

func (apiServer ApiServer) Start() {
	logger.Info("API server started and listening on: ", apiServer.server.Addr)
	apiServer.server.ListenAndServe()
}
