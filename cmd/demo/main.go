package main

import (
	"fmt"
	"strconv"

	demo "github.com/itanoJ/blockchain-go-study"
)

func main() {
	chain := demo.NewGenesisBlockChain("New Genesis Block Chain!")
	chain.AddBlock("A Send to B 1 Btc")
	chain.AddBlock("B Send to C 0.3 Btc")
	for i, b := range chain.Blocks {
		fmt.Printf("index: %d\n", i)
		fmt.Printf("  Data: %s\n", string(b.Data))
		fmt.Printf("  PrevHash: %x\n", b.PrevHash)
		fmt.Printf("  Hash: %x\n", b.Hash)
		fmt.Printf("  Timestamp: %d\n", b.Timestamp)

		proof := demo.NewProofWork(b)
		bol := demo.Validate(proof)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(bol))
	}
}
