package main

import (
	"fmt"

	demo "github.com/itanoJ/blockchain-go-study"
)

func main() {
	chain := demo.NewGenesisBlockChain("New Genesis Block Chain!")
	chain.AddBlockChain("A Send to B 1 Btc")
	chain.AddBlockChain("B Send to C 0.3 Btc")
	for i, b := range chain.Blocks {
		fmt.Printf("index: %d\n", i)
		fmt.Printf("  Data: %s\n", string(b.Data))
		fmt.Printf("  PrevHash: %x\n", b.PrevHash)
		fmt.Printf("  Hash: %x\n", b.Hash)
		fmt.Printf("  Timestamp: %d\n", b.Timestamp)
	}
}
