package main

import "bytes"

type TXOutput struct {
	Value int //store the output with TX
	//ScriptPubKey string //use puzzle to lock output, which is stored in it
	PubKeyHash []byte //store the address, which can unlock the output
}

/*func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}*/


//lock the output which only can be unlocked with specified address
func (out *TXOutput) Lock(address []byte) {
	pubKeyHash := Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4] //why? may be relative with base58decode

	out.PubKeyHash = pubKeyHash
}

//check weather the output is locked with this address
func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

//create a new TX output locked by address
func NewTXOutput(value int, address string) *TXOutput {
	txo := TXOutput{value, nil}
	txo.Lock([]byte(address))

	return &txo
}
