package main

import (
	"blockcain/internal/blockchain"
	"blockcain/internal/config"
	"blockcain/internal/database"
	"blockcain/internal/stress"
	"log"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal("failed to connect to database: %v", err)
	}

	err = database.AutoMigrate(db)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// 블록체인 클라이언트 초기화
	client, err := blockchain.NewBlockChainClient(cfg.RPCURL, cfg.PrivateKey, cfg.FromAddress)
	if err != nil {
		log.Fatalf("failed to create blockchain client: %v", err)
	}

	// 스트레스 테스트 설정
	stressConfig := stress.StressTestConfig{
		NumTransaction:      1000,                 // 생성할 트랜잭션 수
		TransactionInterval: 1 * time.Millisecond, // 트랜잭션 생성 간격
	}

	// 스트레스 테스트 실행
	metrics := stress.RunStressTest(client, stressConfig)

	log.Printf("stress test completed. Metrics: %+v", metrics)
}
