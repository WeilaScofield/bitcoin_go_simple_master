package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

type BlockChainIterator struct {
	db          *bolt.DB
	currentHash []byte
}

func (bci *BlockChainIterator) Next() *Block {
	var block *Block

	err := bci.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(bci.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})
	if err != nil {
		fmt.Println("bci.db.view failed:", err)
	}

	bci.currentHash = block.PrevBlockHash

	return block
}
