package stress

import (
	"blockcain/internal/blockchain"
	"log"
	"sync"
	"time"
)

type StressTestConfig struct {
	NumTransaction      int
	TransactionInterval time.Duration
}

type PerformanceMetrics struct {
	TotalTransactionsSent int
	TransactionDuration   time.Duration
}

// RunStressTest 스트레스 테스트 실행
func RunStressTest(client *blockchain.BlockchainClient, config StressTestConfig) PerformanceMetrics {
	var wg sync.WaitGroup
	transactionCh := make(chan struct{}, config.NumTransaction)
	metrics := PerformanceMetrics{}
	start := time.Now()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < config.NumTransaction; i++ {
			go func() {
				toAddress := "0xSomeRandomAddress"
				value := int64(100)
				_, err := client.SendTransaction(toAddress, value)
				if err != nil {
					log.Printf("failed to send transaction: %v", err)
				}
				transactionCh <- struct{}{}
			}()
			time.Sleep(config.TransactionInterval)
		}
		close(transactionCh)
	}()

	for range transactionCh {
		metrics.TotalTransactionsSent++
	}
	metrics.TransactionDuration = time.Since(start)
	log.Println("Stress test completed")
	return metrics
}
