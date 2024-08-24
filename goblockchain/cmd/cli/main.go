package main

import (
	"fmt"
	"goblockchain/models/block"
	"goblockchain/models/wallet"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {

	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)

	b := block.NewBlockchain(walletM.BlockchainAddress(), 5000)
	isAded := b.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0, walletA.PublicKey(), t.GenerateSignature())

	fmt.Println("Added?", isAded)

	b.Mining()
	b.Print()
	fmt.Printf("A %.1f\n", b.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B %.1f\n", b.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("M %.1f\n", b.CalculateTotalAmount(walletM.BlockchainAddress()))
}
