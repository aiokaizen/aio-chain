package blockchain

import (
	"fmt"
	"math/rand"
	"time"
)

type Blockchain struct {
	Name      string
	Chain     []*Block
	PendingTx []*Transaction
}

func CreateBlockchain(name string) *Blockchain {
	bc := &Blockchain{
		Name: name,
	}
	bc.CreateGenessisBlock()
	return bc
}

func (bc *Blockchain) CreateGenessisBlock() *Block {
	data := fmt.Sprintf("{\"name\": \"%v\", \"settings\": {}}", bc.Name)
	block := bc.CreateBlock(data)
	return block
}

func (bc *Blockchain) CreateBlock(data string) *Block {

	var prevBlock *Block
	prevHash := [32]byte{}

	if len(bc.Chain) > 0 {
		prevBlock = bc.Chain[len(bc.Chain)-1]
		prevHash = prevBlock.Hash
	}

	block := &Block{
		Hash:      [32]byte{},
		PrevHash:  prevHash,
		Data:      []byte(data),
		Timestamp: time.Now().UnixMilli(),
		Nonce:     rand.Int63(),
	}
	block.Hash = block.CalculateHash()

	bc.Chain = append(bc.Chain, block)

	return block
}

func (bc *Blockchain) Print() {
	tx := "[]"

	chain := "["
	for i := 0; i < len(bc.Chain); i++ {
		block := bc.Chain[i]
		chain += block.Repr() + ","
	}
	chain += "\n\t]"

	fmt.Printf(
		"\n\n{\n\t\"Name\": \"%v\",\n\t\"Chain\": %v,\n\t\"Pending transactions\": %v\n}\n\n", bc.Name, chain, tx,
	)
}
