package main

import (
	"fmt"
	"goblockchain/models"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	myBlockchainAddres := "my_blockchain_address"
	blockChain := models.NewBlockchain(myBlockchainAddres)
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 1.0)
	blockChain.Mining()
	blockChain.Print()

	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("X", "Y", 3.0)
	blockChain.Mining()
	blockChain.Print()

	fmt.Printf("my %.1f\n", blockChain.CalculateTotalAmount("my_blockchain_address"))
	fmt.Printf("C %.1f\n", blockChain.CalculateTotalAmount("C"))
	fmt.Printf("D %.1f\n", blockChain.CalculateTotalAmount("D"))
}
