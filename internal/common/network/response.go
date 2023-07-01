package network

import (
	"encoding/json"
	"net/http"

	"github.com/jwo_auth/internal/common/logger"
)

func RespondWithResponseModel(model any, resWriter http.ResponseWriter) {
	RespondWithJsonWithCustomStatusCode(model, resWriter, http.StatusOK)
}

func RespondWithJsonWithCustomStatusCode(data any, resWriter http.ResponseWriter, statusCode int) {
	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(statusCode)

	jsonString, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	logger.Info(string(jsonString))
	resWriter.Write(jsonString)
}
