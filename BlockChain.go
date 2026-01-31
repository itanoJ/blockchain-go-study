package blockchain_go_demo

type BlockChain struct {
	Blocks []*Block
}

func NewGenesisBlockChain(data string) *BlockChain {
	chain := BlockChain{}
	ngb := NewGenesisBlock(data)
	chain.Blocks = append(chain.Blocks, ngb)
	return &chain
}

func (bc *BlockChain) AddBlockChain(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	nb := NewBlock(prevBlock.Hash, data)
	bc.Blocks = append(bc.Blocks, nb)
}
