package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/duddyV/grpc-go-course/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{1, 5, 3, 6, 2, 20}

		for _, number := range numbers {
			log.Printf("Sending number: %d\n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err :=  stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Println("Problem while reading server stream: %v", err)
				break
			}

			log.Printf("The maximum is: %d\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
