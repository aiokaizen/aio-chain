package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Hash      [32]byte
	PrevHash  [32]byte
	Data      []byte
	Timestamp int64
	Nonce     int64
}

func (b *Block) CalculateHash() [32]byte {
	data := bytes.Join(
		[][]byte{
			b.Data,
			b.PrevHash[:],
		}, []byte{},
	)
	hash := sha256.Sum256(data)
	return hash
}

func (block *Block) Repr() string {
	return fmt.Sprintf(
		"\n\t\t{\n\t\t\t\"hash\": \"%v\", \n\t\t\t\"prevHash\": \"%v\", \n\t\t\t\"data\": %v, \n\t\t\t\"timestamp\": %v, \n\t\t\t\"nonce\": %v\n\t\t}",
		hex.EncodeToString(block.Hash[:]), hex.EncodeToString(block.PrevHash[:]),
		string(block.Data), block.Timestamp, block.Nonce,
	)
}
