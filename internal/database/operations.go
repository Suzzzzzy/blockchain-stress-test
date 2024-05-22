package database

import "gorm.io/gorm"

func AddBlock(db *gorm.DB, block Block) error {
	result := db.Create(&block)
	return result.Error
}

func AddTransaction(db *gorm.DB, transaction Transaction) error {
	result := db.Create(&transaction)
	return result.Error
}

func GetBlockByNumber(db *gorm.DB, blockNumber uint) (Block, error) {
	var block Block
	result := db.Preload("Transactions").First(&block, "block_number = ?", blockNumber)
	return block, result.Error
}

func GetTransactionByID(db *gorm.DB, transactionID string) (Transaction, error) {
	var transaction Transaction
	result := db.First(&transaction, "transaction_id = ?", transactionID)
	return transaction, result.Error
}
