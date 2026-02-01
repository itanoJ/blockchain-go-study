package blockchain_go_demo

import (
	"bytes"
	"encoding/gob"
	"time"
)

type Block struct {
	Data      []byte
	PrevHash  []byte
	Hash      []byte
	Timestamp int64
	Nonce     int64
}

func NewBlock(data string, prevHash []byte) *Block {
	block := Block{Data: []byte(data), PrevHash: prevHash, Timestamp: time.Now().UnixNano()}
	proof := NewProofWork(&block)
	nonce, hash := proof.Pow()
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

func NewGenesisBlock(data string) *Block {
	return NewBlock(data, []byte{})
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		panic(err)
	}

	return result.Bytes()
}
