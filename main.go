package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	Blocks []*Block
}
type Block struct {
	timestamp   string
	data        []byte
	none        int
	index       uint
	blockNumber uint
	hash        []byte
	prevHash    []byte
}

func tempDerivHash(b *Block) {
	input := bytes.Join([][]byte{b.data, b.prevHash, []byte(b.timestamp)}, []byte{})
	hash := sha256.Sum256(input)
	b.hash = hash[:]
}
func main() {
	var block Block
	_ = block
	var chain BlockChain
	block.timestamp = "a random timestamp"
	block.data = []byte(`{
        "Name": "Deeksha",  
        "Address": "Hyderabad",
        "Age": 22,
		"n"
    }`)
	chain.Blocks = append(chain.Blocks, &block)
	tempDerivHash(&block)
	fmt.Printf("%#x", block.hash)
	fmt.Printf("\n")
	var block1 Block
	block1.timestamp = "a random timestamp"
	block1.data = []byte(`{
        "Name": "Deeksha",  
        "Address": "Hyderabad",
        "Age": 21
    }`)
	block1.prevHash = chain.Blocks[len(chain.Blocks)-1].hash
	tempDerivHash(&block1)
	chain.Blocks = append(chain.Blocks, &block1)

	fmt.Printf("%#x", block1.hash)
}
