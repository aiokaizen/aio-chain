package blockchain

import (
	"bytes"
	"crypto/sha256"

	"github.com/aiokaizen/aio-chain/tools"
)

type Transaction struct {
	Hash    [32]byte
	From    [32]byte
	To      [32]byte
	Ammount float64
}

func (tx *Transaction) CalculateHash() [32]byte {
	data := bytes.Join(
		[][]byte{tx.From[:], tx.To[:], tools.Float64ToBytes(tx.Ammount)},
		[]byte{},
	)
	hash := sha256.Sum256(data)
	return hash
}

func CreateTransaction(from [32]byte, to [32]byte, ammount float64) *Transaction {
	tx := &Transaction{
		Hash:    [32]byte{},
		From:    from,
		To:      to,
		Ammount: ammount,
	}
	tx.Hash = tx.CalculateHash()
	return tx
}
