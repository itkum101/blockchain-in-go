package main

type Blockchain struct {
	blocks []Block
}

func (b *Blockchain) AddBlock(data string) {
	prevBlock := b.blocks[len(b.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	b.blocks = append(b.blocks, *newBlock)

}

func NewBlockChain() Blockchain {
	return Blockchain{[]Block{*NewGensisBlock()}}
}
