package data

import (
	"context"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/rpc"
)

type Models struct {
	Balance Balance
}

type Balance struct {
	Account string `bson:"account" json:"account"`
}

func (l *Balance) GetBalance(b Balance) (uint64, error) {

	client := client.NewClient(rpc.DevnetRPCEndpoint)
	balance, err := client.GetBalanceWithConfig(
		context.TODO(), 												// request context
		b.Account, 												// wallet to fetch balance for
		rpc.GetBalanceConfig{
			Commitment: rpc.CommitmentProcessed,
		},
	)
	if err != nil {
		return 0, err
	}
	
	return  balance, nil
}