package database

import (
	"gorm.io/gorm"
	"time"
)

type Block struct {
	gorm.Model
	BlockNumber  uint          `gorm:"uniqueIndex"` // 블록 번호, 고유 인덱스
	PreviousHash string        // 이전 블록의 해시
	Timestamp    time.Time     // 블록 생성 시간
	Transactions []Transaction `gorm:"foreignKey:BlockID"` // 블록에 포함된 트랜잭션 목록
}

type Transaction struct {
	gorm.Model
	TransactionID string    `gorm:"uniqueIndex"` // 트랜잭션 ID, 고유 인덱스
	BlockID       uint      // 트랜잭션이 포함된 블록의 ID
	FromAddress   string    // 송신 주소
	ToAddress     string    // 수신 주소
	Amount        float64   // 송신 금액
	Timestamp     time.Time // 트랜잭션 생성 시간
}
