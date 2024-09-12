package main

import (
	"io"
	"log"

	pb "github.com/duddyV/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Avg(stream grpc.ClientStreamingServer[pb.AvgRequest, pb.AvgResponse]) error {
	log.Println("Avg function was invoked")

	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse {
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving req: %v\n", req)

		sum += req.Number
		count++
	}
}
