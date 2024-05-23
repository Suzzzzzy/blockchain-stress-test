package main

import (
	"blockchain/internal/blockchain"
	"blockchain/internal/database"
	"blockchain/internal/util"
)

func main() {
	logger := util.NewLogManager()

	db, err := database.Connect()
	if err != nil {
		logger.Error("failed to connect to database: " + err.Error())
	}

	err = database.AutoMigrate(db)
	if err != nil {
		logger.Error("failed to migrate database: " + err.Error())
	} else {
		logger.Info("database migration completed successfully")
	}

	// 블록체인 생성
	bc := blockchain.NewBlockchain()
	// 트랜잭션 생성
	tx1 := blockchain.NewTransaction("addr1", "add2", 10.0)
	tx2 := blockchain.NewTransaction("addr3", "addr4", 20.0)
	// 블록 생성 및 추가
	block := blockchain.NewBlock("0", []*database.Transaction{tx1, tx2})
	bc.AddBlock(block)
	// 데이터베이스에 블록 저장
	dbHandler := database.NewDatabaseManager(db)
	err = dbHandler.AddBlock(block)
	if err != nil {
		logger.Error("failed to save block to db: " + err.Error())
	}
}
