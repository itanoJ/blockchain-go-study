package main

import "testing"

func TestNewBlockchain_HasGenesisBlock(t *testing.T) {
	bc := NewBlockchain()
	if bc == nil {
		t.Fatal("NewBlockchain() returned nil")
	}
	if len(bc.blocks) != 1 {
		t.Fatalf("expected 1 block (genesis), got %d", len(bc.blocks))
	}
	g := bc.blocks[0]
	if len(g.Hash) == 0 {
		t.Fatal("genesis block hash should not be empty")
	}
	if len(g.PrevBlockHash) != 0 {
		t.Fatalf("genesis prev hash should be empty, got %x", g.PrevBlockHash)
	}
}

func TestAddBlock_LinksPrevHash(t *testing.T) {
	bc := NewBlockchain()
	bc.AddBlock("tx1")
	bc.AddBlock("tx2")

	if len(bc.blocks) != 3 {
		t.Fatalf("expected 3 blocks, got %d", len(bc.blocks))
	}

	b1 := bc.blocks[1]
	b2 := bc.blocks[2]

	if string(b2.PrevBlockHash) != string(b1.Hash) {
		t.Fatalf("expected b2.PrevBlockHash == b1.Hash, got %x vs %x", b2.PrevBlockHash, b1.Hash)
	}
	if len(b1.Hash) == 0 || len(b2.Hash) == 0 {
		t.Fatal("block hash should not be empty")
	}
}
