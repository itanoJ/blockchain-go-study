package blockchain_go_demo

import "github.com/boltdb/bolt"

type BlockChainInterate struct {
	Current []byte
	db      *bolt.DB
}

func (bc *BlockChain) Iterator() *BlockChainInterate {
	bci := &BlockChainInterate{Current: bc.Tip, db: bc.Db}
	return bci
}

func (bc *BlockChainInterate) Next() *Block {
	var bl *Block
	if len(bc.Current) == 0 {
		return nil
	}
	bc.db.View(func(tx *bolt.Tx) error {
		buc := tx.Bucket([]byte(bucketName))
		res := buc.Get(bc.Current)
		if res == nil {
			return nil
		}
		bl = DeserializeBlock(res)
		return nil
	})
	bc.Current = bl.PrevHash
	return bl
}
