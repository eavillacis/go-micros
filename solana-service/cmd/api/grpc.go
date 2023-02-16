package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"solana-service/data"
	"solana-service/solana"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


type SolanaServer struct {
	solana.UnimplementedSolanaServiceServer
	Models data.Models
}

func (s *SolanaServer) GetBalance(ctx context.Context, req *solana.BalanceRequest) (*solana.BalanceResponse, error) {
	input := req.GetAccount()

	// write the log
	balance := data.Balance {
		Account: input.Balance,
	}

	b, err := s.Models.Balance.GetBalance(balance)
	if err != nil {
		res := &solana.BalanceResponse{
			Balance: 0,
			Error: "failed",
		}
		return res, err
	}

	// return response
	res := &solana.BalanceResponse{
		Balance: b,
		Error: "nil",
	}
	return res, nil
}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	// Register reflection service on gRPC server.
	reflection.Register(s)

	solana.RegisterSolanaServiceServer(s, &SolanaServer{Models: app.Models})

	log.Printf("gRPC Server started on port %s", gRpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}