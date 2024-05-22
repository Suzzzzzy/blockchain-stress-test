package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

type BlockchainClient struct {
	client *ethclient.Client
	auth   *bind.TransactOpts
	from   string
}

// NewBlockChainClient 블록체인 client 생성
func NewBlockChainClient(rpcURL, privateKey, fromAddress string) (*BlockchainClient, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(privateKey), "", big.NewInt(1))
	if err != nil {
		return nil, err
	}
	return &BlockchainClient{
		client: client,
		auth:   auth,
		from:   fromAddress,
	}, nil
}

func (bc *BlockchainClient) SendTransaction(toAddress string, value int64) (*types.Transaction, error) {
	nonce, err := bc.client.PendingNonceAt(context.Background(), bc.auth.From)
	if err != nil {
		return nil, err
	}

	gasPrice, err := bc.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), big.NewInt(value), uint64(2100), gasPrice, nil)

	signedTx, err := bc.auth.Signer(bc.auth.From, tx)
	if err != nil {
		return nil, err
	}
	err = bc.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}
