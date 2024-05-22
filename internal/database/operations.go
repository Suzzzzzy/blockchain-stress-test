package database

import "gorm.io/gorm"

type DatabaseManager struct {
	db *gorm.DB
}

func NewDatabaseManager(db *gorm.DB) *DatabaseManager {
	return &DatabaseManager{db: db}
}

func (dm *DatabaseManager) AddBlock(block *Block) error {
	result := db.Create(block)
	return result.Error
}

func (dm *DatabaseManager) AddTransaction(transaction *Transaction) error {
	return dm.db.Create(&transaction).Error

}

func (dm *DatabaseManager) GetBlockByNumber(blockNumber uint) (*Block, error) {
	var block Block
	err := dm.db.Where("block_number = ?", blockNumber).First(&block).Error
	if err != nil {
		return nil, err
	}
	return &block, err
}

func (dm *DatabaseManager) GetTransactionByID(transactionID string) (*Transaction, error) {
	var transaction Transaction
	result := db.First(&transaction, "transaction_id = ?", transactionID)
	return &transaction, result.Error
}
