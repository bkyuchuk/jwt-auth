package network

import (
	"encoding/json"
	"net/http"
)

func ParseRequestTo[T any](req *http.Request) *T {
	var model T
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(model)

	if err != nil {
		panic("Decoding JSON failed with: " + err.Error())
	}
	return &model
}
