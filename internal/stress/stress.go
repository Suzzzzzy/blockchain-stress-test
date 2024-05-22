package stress

import (
	"blockchain/internal/database"
	"blockchain/internal/util"
	"log"
	"strconv"
	"sync"
	"time"
)

// StressTestConfig 스트레스 테스트 설정
type StressTestConfig struct {
	NumBlocks           int           // 생성할 블록 수
	NumTransactions     int           // 생성할 트랜잭션 수
	BlockInterval       time.Duration // 각 블록 생성 간격
	TransactionInterval time.Duration // 각 트랜잭션 생성 간격
}

// PerformanceMetrics 성능 지표 구조체
type PerformanceMetrics struct {
	TotalBlocksCreated          int           // 생성된 총 블록 수
	TotalTransactionsCreated    int           // 생성된 총 트랜잭션 수
	BlockCreationDuration       time.Duration // 블록 생성에 소요된 총 시간
	TransactionCreationDuration time.Duration // 트랜잭션 생성에 소요된 총 시간
}

// RunStressTest 스트레스 테스트 실행 함수
func RunStressTest(db *database.DatabaseManager, config StressTestConfig) PerformanceMetrics {
	logger := util.NewLogManager()

	var wg sync.WaitGroup
	blockCh := make(chan *database.Block, config.NumBlocks)
	transactionCh := make(chan *database.Transaction, config.NumTransactions)

	metrics := PerformanceMetrics{}
	start := time.Now()

	// 고루틴을 사용하여 블록 생성
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < config.NumBlocks; i++ {
			block := &database.Block{
				ID:           strconv.Itoa(i + 1),
				PreviousHash: "0", // 이전 해시값을 임시로 0으로 설정
				Timestamp:    time.Now(),
			}
			time.Sleep(config.BlockInterval)
			blockCh <- block
		}
		close(blockCh)
	}()

	// 고루틴을 사용하여 트랜잭션 생성
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < config.NumTransactions; i++ {
			transaction := &database.Transaction{
				ID:          strconv.Itoa(i + 1),
				FromAddress: "Sender",
				ToAddress:   "Receiver",
				Amount:      10.0,
				Timestamp:   time.Now(),
			}
			time.Sleep(config.TransactionInterval)
			transactionCh <- transaction
		}
		close(transactionCh)
	}()

	// 블록을 데이터베이스에 저장
	go func() {
		for block := range blockCh {
			startBlock := time.Now()
			err := db.AddBlock(block)
			if err != nil {
				logger.Error("failed to add block: " + err.Error())
			}
			metrics.TotalBlocksCreated++
			metrics.BlockCreationDuration += time.Since(startBlock)
		}
	}()

	// 트랜잭션을 데이터베이스에 저장
	go func() {
		for transaction := range transactionCh {
			startTransaction := time.Now()
			err := db.AddTransaction(transaction)
			if err != nil {
				logger.Error("failed to add transaction: " + err.Error())
			}
			metrics.TotalTransactionsCreated++
			metrics.TransactionCreationDuration += time.Since(startTransaction)
		}
	}()

	wg.Wait()

	// 전체 실행 시간 계산
	metrics.BlockCreationDuration = time.Since(start)
	metrics.TransactionCreationDuration = time.Since(start)

	log.Println("Stress test completed")
	return metrics
}
