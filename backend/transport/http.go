package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"depixen/backend/endpoint" 
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrBadRequest = errors.New("bad request")
)


func NewHTTPHandler(endpoints endpoint.Endpoints, router *mux.Router) http.Handler {
	
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(nil)), 
	}

	
	router.Handle("/cards", httptransport.NewServer(
		endpoints.CreateCardEndpoint,
		decodeCreateCardRequest,
		encodeResponse,
		options...,
	)).Methods("POST")

	
	router.Handle("/cards", httptransport.NewServer(
		endpoints.ListCardsEndpoint,
		decodeListCardsRequest,
		encodeResponse,
		options...,
	)).Methods("GET")

	return router
}



func decodeCreateCardRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.CreateCardRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrBadRequest
	}
	return req, nil
}

func decodeListCardsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return endpoint.ListCardsRequest{}, nil
}



func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(endpoint.CreateCardResponse); ok && e.Error != "" {
		encodeError(context.Background(), errors.New(e.Error), w)
		return nil
	}
	if e, ok := response.(endpoint.ListCardsResponse); ok && e.Error != "" {
		encodeError(context.Background(), errors.New(e.Error), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}


func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrBadRequest:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}