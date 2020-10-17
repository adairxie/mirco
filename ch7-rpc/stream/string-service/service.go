package service

import (
	"context"
	"errors"
	"io"
	"log"
	stream_pb "rpc/stream-pb"
	"strings"
)

const (
	StrMaxSize = 1024
)

// Serivce errors
var (
	ErrMaxSize = errors.New("maxmimum size of 1024 bytes exceeded")

	ErrStrValue = errors.New("maximum size of 1024 bytes exceeded")
)

// StringService implement string service
type StringService struct{}

// LotsOfServerStream server stream response
func (s *StringService) LotsOfServerStream(req *stream_pb.StringRequest, qs stream_pb.StringService_LotsOfServerStreamServer) error {
	response := stream_pb.StringResponse{Ret: req.A + req.B}
	for i := 0; i < 10; i++ {
		qs.Send(&response)
	}
	return nil
}

// LotsOfClientStream implement client stream
func (s *StringService) LotsOfClientStream(qs stream_pb.StringService_LotsOfClientStreamServer) error {
	var params []string
	for {
		in, err := qs.Recv()
		if err == io.EOF {
			qs.SendAndClose(&stream_pb.StringResponse{Ret: strings.Join(params, "")})
			return nil
		}
		if err != nil {
			log.Printf("failed to recv: %v", err)
			return err
		}
		params = append(params, in.A, in.B)
	}
}

// LotsOfServerAndClientStream implement client and server stream
func (s *StringService) LotsOfServerAndClientStream(qs stream_pb.StringService_LotsOfServerAndClientStreamServer) error {
	for {
		in, err := qs.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("failed to recv %v", err)
			return err
		}
		qs.Send(&stream_pb.StringResponse{Ret: in.A + in.B})
	}
}

// Concat implements concat interface
func (s *StringService) Concat(ctx context.Context, req *stream_pb.StringRequest) (*stream_pb.StringResponse, error) {
	if len(req.A)+len(req.B) > StrMaxSize {
		response := stream_pb.StringResponse{Ret: ""}
		return &response, nil
	}
	response := stream_pb.StringResponse{Ret: req.A + req.B}
	return &response, nil
}
