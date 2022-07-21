package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"proto"
	"strconv"
	"time"

	"blockchain"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type server struct {
	proto.ClientServerServiceServer
}

const (
	BTC = iota
	ETH
	ADA
)

var bcBTC blockchain.Blockchain
var bcETH blockchain.Blockchain
var bcADA blockchain.Blockchain

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
	case BTC:
		bcBTC = blockchain.CreateBlockchain(int(difficulty))
	case ETH:
		bcETH = blockchain.CreateBlockchain(int(difficulty))
	case ADA:
		bcADA = blockchain.CreateBlockchain(int(difficulty))
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

		from := req.GetBlock().GetFrom()
		to := req.GetBlock().GetTo()
		amount := req.GetBlock().GetAmount()
		token := req.GetToken()

		switch token {
		case BTC:
			bcBTC.AddBlock(from, to, amount)
		case ETH:
			bcETH.AddBlock(from, to, amount)
		case ADA:
			bcADA.AddBlock(from, to, amount)
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
		case BTC:
			if bcBTC.IsValid() == true {
				result = "Check #" + strconv.Itoa(i) + " is valid."
			} else {
				result = "Check #" + strconv.Itoa(i) + " is invalid."
			}
		case ETH:
			if bcETH.IsValid() == true {
				result = "Check #" + strconv.Itoa(i) + " is valid."
			} else {
				result = "Check #" + strconv.Itoa(i) + " is invalid."
			}
		case ADA:
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
