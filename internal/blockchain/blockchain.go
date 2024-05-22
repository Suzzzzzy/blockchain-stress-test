package blockchain

import (
	"blockchain/internal/database"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// NewBlock 함수는 새로운 블록을 생성합니다.
func NewBlock(previousHash string, transactions []*database.Transaction) *database.Block {
	merkleRootHash := calculateMerkleRootHash(transactions)
	return &database.Block{
		PreviousHash:   previousHash,
		Timestamp:      time.Now(),
		MerkleRootHash: merkleRootHash,
	}
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
	str := strconv.FormatUint(tx.ID, 10) + strconv.FormatUint(tx.FromUserID, 10) + strconv.FormatUint(tx.ToUserID, 10) + strconv.FormatFloat(tx.Amount, 'f', -1, 64) + tx.Timestamp.String()
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

// buildMerkleTree 함수는 주어진 해시들로부터 머클 루트 해시를 계산합니다.
func buildMerkleTree(hashes []string) string {
	// 여기에서 실제 머클 트리를 구현하는 로직을 추가합니다.
	// 지금은 간단히 첫 번째 해시를 머클 루트 해시로 사용합니다.
	if len(hashes) > 0 {
		return hashes[0]
	}
	return ""
}

// FindTransactionByID 함수는 특정 ID를 가진 트랜잭션을 찾습니다.
func FindTransactionByID(blockchain []*database.Block, transactionID uint64) *database.Transaction {
	// 블록체인의 모든 블록을 반복하면서 트랜잭션을 찾습니다.
	for _, block := range blockchain {
		// 블록에 포함된 모든 트랜잭션을 검색합니다.
		for _, tx := range block.Transaction {
			if tx.ID == transactionID {
				return tx // 트랜잭션을 찾으면 반환합니다.
			}
		}
	}
	return nil // 트랜잭션을 찾지 못하면 nil을 반환합니다.
}
