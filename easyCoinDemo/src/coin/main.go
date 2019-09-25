package main

import (
	"core"
	"fmt"
	"strconv"
)
const targetBits=16
func main(){
	bc:=core.NewBlockchain()

	bc.AddBlock("Send 1 PPC to BOSS.Pei")
	bc.AddBlock("Send 2 more PPC to BOSS.Pei")

	for _,block:=range bc.Blocks{
		fmt.Printf("Index:%d\n",block.Index)
		fmt.Printf("Prev.hash:%x\n",block.PrevBlockHash)
		fmt.Printf("Data:%s\n",block.Data)
		fmt.Printf("Hash:%x\n",block.Hash)

		pow:=core.NewProofOfWork(block)
		fmt.Printf("PoW:%s\n",strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}