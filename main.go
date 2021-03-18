package main

import (
	"log"

	"github.com/iamcmnut/go-simple-blockchain/blockchain"
)

func main() {
	chain := &blockchain.Blockchain{}

	chain.Init()

	chain.CreateBlock("iamcmnut")
	chain.CreateBlock("test")
	chain.CreateBlock("8782323")

	for i, b := range chain.Blocks {
		log.Printf("Block #%03d, Data: %s, PreviousHash: %s, CurrentHash: %s", i, b.Data, b.PreviousBlockHash, b.GetHash())
	}
}
