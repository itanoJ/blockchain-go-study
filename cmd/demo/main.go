package main

import (
	demo "github.com/itanoJ/blockchain-go-study"
)

func main() {
	bc := demo.NewBlockChain()
	defer bc.Db.Close()

	cli := demo.CLI{bc}
	cli.Run()
}
