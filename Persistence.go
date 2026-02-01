package blockchain_go_demo

import (
	"bytes"
	"encoding/gob"
)

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}

	return &block
}
