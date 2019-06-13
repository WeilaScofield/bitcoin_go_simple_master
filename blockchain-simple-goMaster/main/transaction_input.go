package main

import "bytes"

type TXInput struct {
	Txid []byte //store the ID of the pre transaction
	Vout int    //store an index of outputs in the pre transaction
	//ScriptSig string //a script provide answer of output's scriptPubKey
	Signature []byte
	PubKey    []byte
}

/*func (in *TXInput) CanUnlockOutPutWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}*/

//check weather address initialed the TX
func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPublicKey(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}
