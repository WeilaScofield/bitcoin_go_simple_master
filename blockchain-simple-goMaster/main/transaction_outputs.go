package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

type TXOutputs struct {
	Outputs []TXOutput
}

func (outs TXOutputs) Serialize() []byte {
	var encoded bytes.Buffer
	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(outs)
	if err != nil {
		log.Panic(err)
	}
	return encoded.Bytes()
}

func DeserializeOutputs(data []byte)TXOutputs{
	var outputs TXOutputs
	decoder:=gob.NewDecoder(bytes.NewReader(data))
	err:=decoder.Decode(&outputs)
	if err!=nil{
		log.Panic(err)
	}
	return outputs
}
