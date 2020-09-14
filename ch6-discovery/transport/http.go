package transport

import (
	"context"
	"discovery/endpoint"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

func decodeSayHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return endpoint.SayHelloRequest{}, nil
}

func decodeDiscoveryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	serviceName := r.URL.Query().Get("serviceName")
	if serviceName == "" {
		return nil, ErrorBadRequest
	}
	return endpoint.DiscoveryRequest{
		ServiceName: serviceName,
	}, nil
}

func decodeHealthCheckRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return endpoint.HealthRequest{}, nil
}

func encodeJsonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
