package main

import (
	"context"
	"strconv"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/rpc"
)

// RPCServer is the type for our RPC Server. Methods that take this as a receiver are available
// over RPC, as long as they are exported.
type RPCServer struct{}

// RPCPayload is the type for data we receive from RPC
type RPCPayload struct {
	Account string
}

func (r *RPCServer) GetBalanceWithConfig(payload RPCPayload, resp *string) error {
	client := client.NewClient(rpc.DevnetRPCEndpoint)
	balance, err := client.GetBalanceWithConfig(
		context.TODO(), 												// request context
		payload.Account, 												// wallet to fetch balance for
		rpc.GetBalanceConfig{
			Commitment: rpc.CommitmentProcessed,
		},
	)
	if err != nil {
		return err
	}

	// resp is the message sent back to the RPC caller
	*resp = "Processed payload via RPC - Balance: " + strconv.FormatUint(balance, 10)
	return nil
}
