package blockchain_go_demo

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Data      []byte
	PrevHash  []byte
	Hash      []byte
	Timestamp int64
}

func (b *Block) SetHash() {
	res := bytes.Join([][]byte{
		b.Data,
		b.PrevHash,
		[]byte(strconv.FormatInt(b.Timestamp, 10)),
	}, []byte{})
	hash := sha256.Sum256(res)
	b.Hash = hash[:]
}

func NewBlock(prevHash []byte, data string) *Block {
	block := Block{Data: []byte(data), PrevHash: prevHash, Timestamp: time.Now().UnixNano()}
	block.SetHash()
	return &block
}

func NewGenesisBlock(data string) *Block {
	genesis := Block{Data: []byte(data), Timestamp: time.Now().UnixNano()}
	genesis.SetHash()
	return &genesis
}
