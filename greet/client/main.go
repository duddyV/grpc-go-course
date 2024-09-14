package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/duddyV/grpc-go-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := true // change that to false if needed
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificateL: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	// conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	// doGreetWithDeadline(c, 5*time.Second)
	// doGreetWithDeadline(c, 1*time.Second)
}
