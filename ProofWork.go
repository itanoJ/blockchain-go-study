package blockchain_go_demo

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

/**
  工作量证明

包含 区块 目标阈值
*/

// 目标阈值
const targetBits = 4

const maxVa = math.MaxInt64

type ProofWork struct {
	Block  *Block
	Target *big.Int
}

/*
*
思路
00001 -》 256 位是最小的值  根据bitcion的原理 将0左移位越小 值就会越小  mining的难度就会越大
因为如果最高位为10000 -》 256   这样的话是最大值，如果想找到比他小的值则会很容易

当我们去构建一个工作量证明的结构体时  我们要将这个target计算出来 目前这个是固定的 但是在bitcion中 这个是动态
调整的
*/
func NewProofWork(block *Block) *ProofWork {
	newInt := big.NewInt(1) // 这个是上界 最大的值
	target := newInt.Lsh(newInt, 256-targetBits)
	return &ProofWork{Block: block, Target: target}
}

/*
	引入了工作量证明 就要明确 原理

目的 找到nonce值
*/
func (p *ProofWork) PrepareData(nonce int64) []byte {
	data := bytes.Join([][]byte{
		p.Block.Data,
		p.Block.PrevHash,
		IntToHex(p.Block.Timestamp),
		IntToHex(nonce),
		IntToHex(targetBits),
	}, []byte{})
	return data
}

/*
*

	这个是去做 pow  工作量证明的地方

这里 回去找nonce值  只有找到了一个合法的nonce值 才能让 当前的 hash结果转换成int
小于target  这个就是工作量证明 只有找到的人 才能享有 往区块链上面增加block的能力
*/
func (p *ProofWork) Pow() (int64, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Println("挖取区块")
	// 无限循环去找一个合适的nonce值
	for nonce < maxVa {
		data := p.PrepareData(int64(nonce))
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(p.Target) == -1 {
			fmt.Println("区块挖取成功，合法的nonce：", nonce)
			break
		}

		nonce++
	}

	return int64(nonce), hash[:]
}

/*
*
将int 64 转换为 []byte切片
*/
func IntToHex(v int64) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, v) // 大端序  区块链中多采用这种方式
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func Validate(p *ProofWork) bool {
	var hashInt big.Int
	r := p.PrepareData(p.Block.Nonce)
	hash := sha256.Sum256(r)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(p.Target) == -1

}
