package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints structure
type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

// MakeEndpoints method
func MakeEndpoints(service Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(service),
		GetUser:    makeGetUserEndpoint(service),
	}
}

func makeCreateUserEndpoint(service Service) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		err := service.CreateUser(context, req.Email, req.Password)
		return CreateUserResponse{}, err
	}
}

func makeGetUserEndpoint(service Service) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		user, err := service.GetUser(context, req.ID)
		return GetUserResponse{User: user}, err
	}
}
