package database

import (
	"time"
)

type Block struct {
	// 블록의 고유 식별자
	ID string
	// 블록의 이전 해시값
	PreviousHash string
	// 블록의 생성 시간
	Timestamp time.Time
	// 블록에 포함된 트랜잭션들의 머클 루트 해시값
	MerkleRootHash string
	Transaction    []*Transaction
}

type Transaction struct {
	// 트랜잭션의 고유 식별자
	ID          string
	FromAddress string
	ToAddress   string
	// 트랜잭션의 금액
	Amount float64
	// 트랜잭션의 생성 시간
	Timestamp time.Time
}

type Blockchain struct {
	Block []*Block
}

// MerkleTree 해시트리
type MerkleTree struct {
	RootHash string
}
