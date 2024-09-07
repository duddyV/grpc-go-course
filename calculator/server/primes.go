package main

import (
	"log"

	pb "github.com/duddyV/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Primes(req *pb.PrimesRequest, stream grpc.ServerStreamingServer[pb.PrimesResponse]) error {
	log.Printf("GreetManyTimes function was invoked with: %v\n", req)

	number := req.Number
	divisor := int64(2)

	for number > 1 {
		if number % divisor == 0 {
			stream.Send(&pb.PrimesResponse{
				Result: divisor,
			})

			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}
