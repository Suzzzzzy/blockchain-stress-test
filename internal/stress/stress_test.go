package stress

import (
	"blockchain/internal/database"
	"log"
	"testing"
	"time"
)

// 스트레스 테스트 실행을 테스트하는 함수
func TestRunStressTest(t *testing.T) {
	db := &database.DatabaseManager{}

	stressConfig := StressTestConfig{
		NumBlocks:           10,                    // 테스트를 위한 블록 수
		NumTransactions:     100,                   // 테스트를 위한 트랜잭션 수
		BlockInterval:       10 * time.Millisecond, // 블록 생성 간격
		TransactionInterval: 1 * time.Millisecond,  // 트랜잭션 생성 간격
	}

	metrics := RunStressTest(db, stressConfig)

	// 블록 및 트랜잭션 수 확인
	var blockCount int64
	db.Model(&database.Block{}).Count(&blockCount)
	if blockCount != int64(stressConfig.NumBlocks) {
		t.Errorf("expected %d blocks, got %d", stressConfig.NumBlocks, blockCount)
	}

	var transactionCount int64
	db.Model(&database.Transaction{}).Count(&transactionCount)
	if transactionCount != int64(stressConfig.NumTransactions) {
		t.Errorf("expected %d transactions, got %d", stressConfig.NumTransactions, transactionCount)
	}

	log.Printf("Stress test passed. Metrics: %+v", metrics)
}
