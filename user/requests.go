package user

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type (
	// CreateUserRequest structure
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// GetUserRequest structure
	GetUserRequest struct {
		ID int `json:"id"`
	}
)

func decodeCreateUserRequest(context context.Context, request *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(request.Body).Decode(&req)
	return req, err
}

func decodeGetUserRequest(context context.Context, request *http.Request) (interface{}, error) {
	vars := mux.Vars(request)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		return nil, err
	}

	req := GetUserRequest{
		ID: id,
	}

	return req, nil
}
