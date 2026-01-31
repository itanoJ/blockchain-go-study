package blockchain_go_demo

type BlockChain struct {
	Blocks []*Block
}

func NewGenesisBlockChain(data string) *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock(data)}}
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	nb := NewBlock(prevBlock.Hash, data)
	bc.Blocks = append(bc.Blocks, nb)
}
