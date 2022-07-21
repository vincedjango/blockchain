package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type Token int32

const (
	BTC Token = iota
	ADA
)

func main() {
	certFile := "../../ssl/ca.crt"
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	if sslErr != nil {
		log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
		return
	}

	opts := grpc.WithTransportCredentials(creds)
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Could not connect to %v", err)
	}
	defer cc.Close()

	client := proto.NewClientServerServiceClient(cc)

	createBlockchain(client, BTC)
	createBlockchain(client, ADA)
	AddBlock(client, BTC)
	AddBlock(client, ADA)
	isValid(client, BTC)
	isValid(client, ADA)
}

// Unary function
func createBlockchain(client proto.ClientServerServiceClient, tok Token) {
	fmt.Println("Create Blockchain")

	in := proto.CreationRequest{
		Difficulty: 3,
		Token:      int32(tok),
	}

	response, err := client.CreateBlockchain(context.Background(), &in)

	if err != nil {
		log.Fatalf("Error while calling client.CreateBlockchain with error: %v", err)
		return
	}

	fmt.Printf("Response from createBlockchain: %v", response.Result)
}

// Client Streaming
func AddBlock(client proto.ClientServerServiceClient, tok Token) {
	fmt.Println("Add Block")

	requests := []*proto.AddBlockRequest{
		&proto.AddBlockRequest{
			Block: &proto.Block{
				From:   "User1",
				To:     "User2",
				Amount: 1,
			},
			Token:  int32(tok),
			Number: 3,
		},
		&proto.AddBlockRequest{
			Block: &proto.Block{
				From:   "User2",
				To:     "User3",
				Amount: 2,
			},
			Token:  int32(tok),
			Number: 1,
		},
		&proto.AddBlockRequest{
			Block: &proto.Block{
				From:   "User3",
				To:     "User4",
				Amount: 3,
			},
			Token:  int32(tok),
			Number: 20,
		},
	}

	stream, err := client.AddBlock(context.Background())
	if err != nil {
		statusErr, ok := status.FromError(err)
		if !ok {
			if statusErr.Code() == codes.Internal {
				log.Fatalf("Internal Error! The block was not created.")
			}
		}
		return
	}

	for _, request := range requests {
		fmt.Printf("Sending request %v \n", request)
		stream.Send(request)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while calling Client CloseAndRecv with error: %v", err)
	}

	fmt.Printf("Add Block response: %v", res)
}

// Server Streaming
func isValid(client proto.ClientServerServiceClient, tok Token) {
	fmt.Println("Is Valid")

	req := &proto.IsValidRequest{
		Redundancy: 2,
		Token:      int32(tok),
	}

	resStream, err := client.IsValid(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling client.IsValid with error: %v", err)
		return
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we have reached the end of the file
			break
		}
		if err != nil {
			log.Fatalf("Client - error while reading stream %v", err)
		}
		log.Printf(msg.GetResult())
	}
}
