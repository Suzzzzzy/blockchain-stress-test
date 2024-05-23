package stress

import (
	"blockchain/internal/database"
	"testing"
	"time"
)

func TestRunStressTest(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Errorf("failed to connect db")
	}
	con := database.NewDatabaseManager(db)

	if err != nil {
		t.Fatalf("Failed to migrate database schema: %v", err)
	}

	// 스트레스 테스트 설정
	config := StressTestConfig{
		NumBlocks:           10,
		NumTransactions:     100,
		BlockInterval:       time.Millisecond * 100,
		TransactionInterval: time.Millisecond * 50,
	}

	metrics := RunStressTest(con, config)

	// 결과 검증
	if metrics.TotalBlocksCreated != config.NumBlocks {
		t.Errorf("Expected %d blocks to be created, but got %d", config.NumBlocks, metrics.TotalBlocksCreated)
	}

	if metrics.TotalTransactionsCreated != config.NumTransactions {
		t.Errorf("Expected %d transactions to be created, but got %d", config.NumTransactions, metrics.TotalTransactionsCreated)
	}

	// 데이터베이스에 저장된 블록과 트랜잭션 수 확인
	var blocksCount int64
	db.Model(&database.Block{}).Count(&blocksCount)
	if blocksCount != int64(config.NumBlocks) {
		t.Errorf("Expected %d blocks to be stored in database, but got %d", config.NumBlocks, blocksCount)
	}

	var transactionsCount int64
	db.Model(&database.Transaction{}).Count(&transactionsCount)
	if transactionsCount != int64(config.NumTransactions) {
		t.Errorf("Expected %d transactions to be stored in database, but got %d", config.NumTransactions, transactionsCount)
	}
}
