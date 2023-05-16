package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 22

// INT TO HEX

// Block
// Store valuable information

// Block struct
// Simple implementation

// Setting a hash for a certain block

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	return &ProofOfWork{block: b, target: target}
}

// DATA TO HASH
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data,
		IntToHex(pow.block.Timestamp),
		IntToHex(int64(nonce)),
		IntToHex(int64(targetBits)),
	}, []byte{})
	return data

}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int

	var hash [32]byte

	nonce := 0

	fmt.Println()

	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)

		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce, hash[:]

}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

func (block *Block) SetHash() {
	// CREATING TIMESTAMP FOR THAT

	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))

	headers := bytes.Join([][]byte{block.Data, block.PrevBlockHash, timestamp}, []byte{})

	hash := sha256.Sum256(headers)

	block.Hash = hash[:]
}

// Creating a new block
// func CreateBlock(data string, prevBlockHash []byte) Block {
// 	block := Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
// 	block.SetHash()
// 	return block
// }

// Adding blocks
