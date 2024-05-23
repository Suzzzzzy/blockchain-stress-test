package blockchain

import (
	"blockchain/internal/database"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Blockchain struct {
	Blocks       []*database.Block
	Transactions map[string]*database.Transaction
}

// MerkleTree 해시트리
type MerkleTree struct {
	RootHash string
}

// NewBlock 함수는 새로운 블록을 생성합니다.
func NewBlock(previousHash string, transactions []*database.Transaction) *database.Block {
	merkleRootHash := calculateMerkleRootHash(transactions)
	return &database.Block{
		ID:             generateBlockID(), // 블록 ID 생성
		PreviousHash:   previousHash,
		Timestamp:      time.Now(),
		MerkleRootHash: merkleRootHash,
		Transactions:   transactions,
	}
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks:       []*database.Block{},
		Transactions: make(map[string]*database.Transaction),
	}
}

func (bc *Blockchain) AddBlock(block *database.Block) {
	bc.Blocks = append(bc.Blocks, block)
	for _, tx := range block.Transactions {
		bc.Transactions[tx.ID] = tx
	}
}

func NewTransaction(from, to string, amount float64) *database.Transaction {
	timestamp := time.Now()
	txID := generateTransactionID(from, to, amount, timestamp)

	return &database.Transaction{
		ID:          txID,
		FromAddress: from,
		ToAddress:   to,
		Amount:      amount,
		Timestamp:   timestamp,
	}
}

func generateTransactionID(from, to string, amount float64, timestamp time.Time) string {
	// 트랜잭션의 고유 ID를 생성하기 위해 송신자, 수신자, 금액, 타임스탬프를 조합합니다.
	data := from + to + strconv.FormatFloat(amount, 'f', -1, 64) + timestamp.String()

	// 생성된 데이터에 대한 SHA-256 해시를 계산하여 반환합니다.
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// calculateMerkleRootHash 함수는 주어진 트랜잭션들을 이용하여 머클 루트 해시를 계산합니다.
func calculateMerkleRootHash(transactions []*database.Transaction) string {
	var hashes []string
	for _, tx := range transactions {
		hash := calculateTransactionHash(tx)
		hashes = append(hashes, hash)
	}
	return buildMerkleTree(hashes)
}

// calculateTransactionHash 함수는 주어진 트랜잭션의 해시를 계산합니다.
func calculateTransactionHash(tx *database.Transaction) string {
	str := tx.ID + tx.FromAddress + tx.ToAddress + strconv.FormatFloat(tx.Amount, 'f', -1, 64) + tx.Timestamp.String()
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

// buildMerkleTree 함수는 주어진 해시들로부터 머클 루트 해시를 계산합니다.
func buildMerkleTree(hashes []string) string {
	// 해시 슬라이스의 길이가 1 이하면 현재 해시를 반환합니다.
	if len(hashes) <= 1 {
		return hashes[0]
	}

	// 새로운 해시 슬라이스를 만들어 두 개씩 연속된 해시를 묶어줍니다.
	var newHashes []string
	for i := 0; i < len(hashes); i += 2 {
		// 마지막 노드는 자신만을 묶어서 처리합니다.
		if i+1 == len(hashes) {
			newHashes = append(newHashes, hashes[i])
		} else {
			// 두 개의 해시를 결합하여 새로운 해시를 생성합니다.
			combinedHash := hashes[i] + hashes[i+1]
			newHashes = append(newHashes, calculateHash(combinedHash))
		}
	}

	// 새로운 해시 슬라이스를 기반으로 재귀적으로 해시 트리를 구축합니다.
	return buildMerkleTree(newHashes)
}

// calculateHash 함수는 주어진 문자열의 해시를 계산합니다.
func calculateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func generateBlockID() string {
	return hex.EncodeToString(sha256.New().Sum(nil))
}
