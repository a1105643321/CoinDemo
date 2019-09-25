package core

import "time"

type Block struct {
	Index int
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
	Nonce int
}

func NewBlock(data string,index int,prevBlockHash []byte)*Block{
	block:=&Block{index,time.Now().Unix(),[]byte(data),prevBlockHash,[]byte{},0}
	var pow = NewProofOfWork(block)
	nonce,hash:=pow.Run()

	block.Hash=hash
	block.Nonce=nonce
	block.Index=index+1

	return block
}

func NewGenisisBlock()*Block{
	return NewBlock("Genisis Block",0,[]byte{})
}