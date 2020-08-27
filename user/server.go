package user

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewServer method
func NewServer(context context.Context, endpoints Endpoints) http.Handler {
	router := mux.NewRouter()

	router.Use(jsonContentTypeMidlleware)

	router.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		decodeCreateUserRequest,
		encodeResponse,
	))

	router.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		decodeGetUserRequest,
		encodeResponse,
	))

	return router
}

func jsonContentTypeMidlleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writter http.ResponseWriter, request *http.Request) {
		writter.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(writter, request)
	})
}
