package blockchain_go_demo

import "github.com/boltdb/bolt"

const bucketName = "bucket"

const dbName = "BlockChain.db"

type BlockChain struct {
	Tip []byte
	Db  *bolt.DB
}

/*
*

	选择在KV数据库中 存放区块链数据
	key                        value
	newGenesisBlock hash key   newGenesisBlock Serialize value
	l                          newGenesisBlock hash key
*/
func NewBlockChain() *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {

		buc := tx.Bucket([]byte(bucketName))
		if buc == nil {
			ng := NewGenesisBlock("New Genesis Block!")
			buck, err := tx.CreateBucket([]byte(bucketName))
			if err != nil {
				panic(err)
			}
			err = buck.Put(ng.Hash, ng.Serialize())
			if err != nil {
				panic(err)
			}
			err = buck.Put([]byte("l"), ng.Hash)
			if err != nil {
				panic(err)
			}
		} else {
			tip = buc.Get([]byte("l"))
		}

		return nil
	})

	bc := BlockChain{Tip: tip, Db: db}
	return &bc

}

func (bc *BlockChain) AddBlock(data string) {
	var latest []byte
	err := bc.Db.View(func(tx *bolt.Tx) error {
		buc := tx.Bucket([]byte(bucketName))
		latest = buc.Get([]byte("l"))
		return nil
	})
	if err != nil {
		panic(err)
	}
	nb := NewBlock(data, latest)
	err = bc.Db.Update(func(tx *bolt.Tx) error {
		buc := tx.Bucket([]byte(bucketName))
		buc.Put(nb.Hash, nb.Serialize())
		buc.Put([]byte("l"), nb.Hash)
		bc.Tip = nb.Hash
		return nil
	})
	if err != nil {
		panic(err)
	}

}
