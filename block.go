package main

import "time"

type Block struct {
	// Timestamp == time when block was created
	Timestamp int64
	// Actual valuable information in block
	Data []byte
	// Stores hash of the previous block
	PrevBlockHash []byte
	// Hash of current block
	Hash []byte

	Nonce int
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func NewGensisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
