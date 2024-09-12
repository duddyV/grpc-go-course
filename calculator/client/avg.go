package main

import (
	"context"
	"log"

	pb "github.com/duddyV/grpc-go-course/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Printf("Error while calling doAvg: %v", err)
	}

	numbers := []int32{1,2,3,4}
	for _, number := range numbers {
		log.Printf("Sending req: %v\n", number)
		stream.Send(&pb.AvgRequest {
			Number: number,
		})
	}

	reqs := []*pb.AvgRequest {
		{ Number: 1 },
		{ Number: 2 },
		{ Number: 3 },
		{ Number: 4 },
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from doAvg: %v\n", err)
	}

	log.Printf("doAvg: %f\n", res.Result)
}
