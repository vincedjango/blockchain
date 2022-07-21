package main

import (
	"blockchain"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"proto"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type server struct {
	proto.ClientServerServiceServer
}

const (
	cBTC = iota
	cADA
)

var bcBTC *blockchain.BlockchainBTC
var bcADA *blockchain.BlockchainADA

func main() {
	// 50051 is the default port for grpc
	fmt.Println("Server Created!")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	certFile := "../../ssl/server.crt"
	keyFile := "../../ssl/server.pem"
	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if sslErr != nil {
		log.Fatalf("Failed loading certificates: %v", sslErr)
		return
	}
	opts := grpc.Creds(creds)

	s := grpc.NewServer(opts)
	proto.RegisterClientServerServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}

// TODO: Having multiple blockchain, so I can add interface
func (*server) CreateBlockchain(ctx context.Context, in *proto.CreationRequest) (*proto.CreationResponse, error) {
	// Greet(ctx context.Context, in *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("CreateBlockchain function was invoked with %v", in)
	difficulty := in.GetDifficulty()
	token := in.GetToken()

	switch token {
	case cBTC:
		bcBTC = blockchain.CreateBlockchainBTC(int(difficulty))
	case cADA:
		bcADA = blockchain.CreateBlockchainADA(int(difficulty))
	}

	response := proto.CreationResponse{
		Result: "The blockchain was created",
	}

	return &response, nil
}

func (*server) AddBlock(stream proto.ClientServerService_AddBlockServer) error {
	fmt.Println("Server AddBlock invoked")

	result := "Block Added."
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// We are done reading the client stream
			return stream.SendAndClose(
				&proto.AddBlockResponse{
					Result: result,
				})
		}
		if err != nil {
			return status.Error(codes.Internal, "The block failed to be created")
		}

		switch req.GetToken() {
		case cBTC:
			blockchain.IAddBlock(bcBTC, req.GetBlock().GetFrom(), req.GetBlock().GetTo(), req.GetBlock().GetAmount(), req.GetNumber())
		case cADA:
			blockchain.IAddBlock(bcADA, req.GetBlock().GetFrom(), req.GetBlock().GetTo(), req.GetBlock().GetAmount(), req.GetNumber())
		}
	}
}

// TODO: Having multiple blockchain, so I can add interface
func (*server) IsValid(req *proto.IsValidRequest, stream proto.ClientServerService_IsValidServer) error {
	// fmt.Println("Server IsValid invoked")
	var result string

	redundancy := req.GetRedundancy()
	token := req.GetToken()

	for i := 0; i < int(redundancy); i++ {

		switch token {
		case cBTC:
			if bcBTC.IsValid() == true {
				result = "Check #" + strconv.Itoa(i) + " is valid."
			} else {
				result = "Check #" + strconv.Itoa(i) + " is invalid."
			}
		case cADA:
			if bcADA.IsValid() == true {
				result = "Check #" + strconv.Itoa(i) + " is valid."
			} else {
				result = "Check #" + strconv.Itoa(i) + " is invalid."
			}
		}

		res := &proto.IsValidResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}
