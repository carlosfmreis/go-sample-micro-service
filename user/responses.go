package user

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	// CreateUserResponse structure
	CreateUserResponse struct{}
	// GetUserResponse structure
	GetUserResponse struct {
		User User `json:"user"`
	}
)

func encodeResponse(context context.Context, writter http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(writter).Encode(response)
}
