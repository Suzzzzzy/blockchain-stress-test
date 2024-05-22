package stress

import (
	"blockcain/internal/blockchain"
	"blockcain/internal/config"
	"log"
	"testing"
	"time"
)

func setupClient(t *testing.T) *blockchain.BlockchainClient {
	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	client, err := blockchain.NewBlockChainClient(cfg.RPCURL, cfg.PrivateKey, cfg.FromAddress)
	if err != nil {
		t.Fatalf("failed to create blockchain client: %v", err)
	}

	return client
}

func TestRunStressTest(t *testing.T) {
	client := setupClient(t)

	stressConfig := StressTestConfig{
		NumTransaction:      100,
		TransactionInterval: 1 * time.Millisecond,
	}
	metrics := RunStressTest(client, stressConfig)
	log.Printf("stress test passed. Metrics: %+v", metrics)
}