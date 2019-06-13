package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"go-ethereum-master/common/math"
	"math/big"
)

const targetBits = 20 //each 4 mean a 0

var maxNonce = math.MaxInt64

type ProofOfWork struct {
	target *big.Int //nonce should less than target
	block  *Block
}



func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	// move left to set bits
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{target, b}
	return pow
}

//content add nonce which need be hash later
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.block.HashTransaction(),
		pow.block.PrevBlockHash,
		IntToHex(pow.block.Timestamp),
		IntToHex(int64(targetBits)),
		IntToHex(int64(nonce)),
	}, []byte{})

	return data
}

//the kernel algorithm of POW
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%v\"\n", pow.block.Transactions)
	for nonce < maxNonce { //protect nonce from overflow
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		//break until present hash is less than target
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Print("\n\n")
	return nonce, hash[:]

}

//verify if the block processed by POW
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
