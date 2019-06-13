package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func IntToHex(i int64) []byte {
	res := fmt.Sprintf("%x", i)
	return []byte(res)
}

func ReadRequest(request []byte) *verzion {
	var buff bytes.Buffer
	var payload verzion

	buff.Write(request[commandLength:])
	decoder:=gob.NewDecoder(&buff)
	err:=decoder.Decode(&payload)
	if err!=nil{
		log.Panic(err)
	}

	return &payload
}

func gobEncode(data interface{}) []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(data)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}