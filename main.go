package main

import (
	"github.com/aiokaizen/aio-chain/blockchain"
)

func main() {
	// tx := blockchain.CreateTransaction(
	// 	[]byte{},
	// 	[]byte{},
	// 	23,
	// )
	bc := blockchain.CreateBlockchain("YTCC")
	bc.CreateBlock("{\"from\": \"Mouad\", \"to\": \"Anas\", \"ammount\": 12.5}")
	bc.CreateBlock("{\"from\": \"Anas\", \"to\": \"Soufiane\", \"ammount\": 4.3}")
	bc.CreateBlock("{\"from\": \"Soufiane\", \"to\": \"Mouad\", \"ammount\": 14.8983745}")

	bc.Print()
}
