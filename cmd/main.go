package main

import (
	"blockchain/internal/config"
	"blockchain/internal/database"
	"blockchain/internal/stress"
	"blockchain/internal/util"
	"log"
	"time"
)

func main() {
	logger := util.NewLogManager()

	_, err := config.LoadConfig()
	if err != nil {
		logger.Error("failed to load config: " + err.Error())
	}

	db, err := database.Connect()
	if err != nil {
		logger.Error("failed to connect to database: " + err.Error())
	}

	err = database.AutoMigrate(db)
	if err != nil {
		logger.Error("failed to migrate database: " + err.Error())
	}

	// 스트레스 테스트 설정
	stressConfig := stress.StressTestConfig{
		NumTransactions:     1000,                 // 생성할 트랜잭션 수
		TransactionInterval: 1 * time.Millisecond, // 트랜잭션 생성 간격
	}

	// 스트레스 테스트 실행
	metrics := stress.RunStressTest(client, stressConfig)

	log.Printf("stress test completed. Metrics: %+v", metrics)
}
