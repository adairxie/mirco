package service

import (
	"context"
	"errors"
	"rpc/pb"
	"strings"
)

// Service constants
const (
	StrMaxSize = 1024
)

// Service errors
var (
	ErrMaxSize  = errors.New("maximum size of 1024 bytes exceeded")
	ErrStrValue = errors.New("maximum size of 1024 bytes exceeded")
)

// StringRequest service request paramas
type StringRequest struct {
	A string
	B string
}

// Service string servic interface
type Service interface {
	// Concat a and b
	Concat(req StringRequest, ret *string) error

	// a, b common string value
}

// StringService  implements Service interface
type StringService struct{}

// Concat concat string a and b
func (s StringService) Concat(ctx context.Context, req *pb.StringRequest) (*pb.StringReponse, error) {
	if len(req.A)+len(req.B) > StrMaxSize {
		response := pb.StringReponse{Ret: ""}
		return &response, nil
	}

	response := pb.StringReponse{Ret: req.A + req.B}
	return &response, nil
}

// Diff common string between a and b
func (s StringService) Diff(ctx context.Context, req *pb.StringRequest) (*pb.StringReponse, error) {
	if len(req.A) < 1 || len(req.B) < 1 {
		response := &pb.StringReponse{Ret: ""}
		return response, nil
	}

	res := ""

	for _, char := range req.A {
		if !strings.Contains(req.B, string(char)) {
			res = res + string(char)
		}
	}

	for _, char := range req.B {
		if !strings.Contains(req.A, string(char)) {
			res = res + string(char)
		}
	}

	response := pb.StringReponse{Ret: res}
	return &response, nil
}
