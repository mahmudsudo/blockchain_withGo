package main

import (
	"blockchain/blocks"
	"fmt"
	"strconv"
)



func main() {
	chain := blocks.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")
	chain.AddBlock("fourth Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blocks.NewProof(block)
		fmt.Printf("pow : %s\n",strconv.FormatBool(pow.Validate()))
	}
}
