package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

/*
  - 定义一个简易的区块结构，参考btc的结构
    我们的结构中不会区分header 和交易体  将两者合并
*/
type Block struct {
	Data          []byte // 存放交易信息
	Hash          []byte // 当前区块的hash = hash(data + prevhash + timestamp)
	PrevBlockHash []byte // 上一个区块的hash
	Timestamp     int64  // 区块生成的时间戳
}

/*
*

	计算当前hash的策略是  通过hash = hash(data + prevhash + timestamp)
*/
func (b *Block) setHash() {
	// 将timestamp转换成字节流
	ts := strconv.FormatInt(int64(b.Timestamp), 10)
	hs := bytes.Join([][]byte{b.Data, b.PrevBlockHash, []byte(ts)}, []byte{})
	res := sha256.Sum256(hs)
	b.Hash = res[:]
}

func NewBlock(data string, prevHash []byte) *Block {
	newBlock := Block{Data: []byte(data), PrevBlockHash: prevHash, Timestamp: time.Now().Unix()}
	newBlock.setHash()
	return &newBlock
}

// 定义一个区块链
type BlockChain struct {
	blocks []*Block
}

/*
*

	将当前区块加入到区块链中
	[]<-[]<-[]<-[]  <--- last block
*/
func (bc *BlockChain) AddBlock(data string) {
	lastBlock := bc.blocks[len(bc.blocks)-1] // 获取最新的去区块
	nBlock := NewBlock(data, lastBlock.Hash)
	bc.blocks = append(bc.blocks, nBlock)
}

/*
*
参考BTC 需要创世纪块
*/
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

/*
*
链上第一个区块 需要我们来创建
*/
func NewBlockchain() *BlockChain {
	return &BlockChain{blocks: []*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
