package blockchain

import (
	"encoding/gob"
	"io"
)

// BlockChain represents a linked list of blocks
type BlockChain struct {
	Blocks []*Block
	Head   Hash
}

// Len returns the length of the BlockChain when marshalled
func (bc *BlockChain) Len() int {
	l := 0
	for _, b := range bc.Blocks {
		l += b.Len()
	}
	return l + HashLen
}

// Marshal converts the BlockChain to a byte slice.
func (bc *BlockChain) Marshal() []byte {
	buf := make([]byte, 0, bc.Len())
	for _, b := range bc.Blocks {
		buf = append(buf, b.Marshal()...)
	}
	return append(buf, bc.Head.Marshal()...)
}

// Encode writes the marshalled blockchain to the given io.Writer
func (bc *BlockChain) Encode(w io.Writer) {
	gob.NewEncoder(w).Encode(bc)
}

// DecodeBlockChain reads the marshalled blockchain from the given io.Reader
func DecodeBlockChain(r io.Reader) *BlockChain {
	var bc BlockChain
	gob.NewDecoder(r).Decode(&bc)
	return &bc
}

// ValidTransaction checks whether a transaction is valid, assuming the blockchain is valid.
func (bc *BlockChain) ValidTransaction(t *Transaction) bool {
	// Find the transaction input (I) in the chain (by hash)
	// Check that output to sender in I is equal to outputs in T
	// Verify signature of T
	return false
}

// ValidBlock checks whether a block is valid
func (bc *BlockChain) ValidBlock(b *Block) bool {
	for _, t := range b.Transactions {
		if !bc.ValidTransaction(t) {
			return false
		}
	}
	// Check that block number is one greater than last block
	// Check that hash of last block is correct
	return false
}
