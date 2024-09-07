package main

import (
	"context"
	"log"

	pb "github.com/duddyV/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", req)

	return &pb.SumResponse{
		Sum: req.A + req.B,
	}, nil
}
