package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	chanvalue                  float32
}

func NewTransaction(sender string, recipent string, value float32) *Transaction {
	return &Transaction{sender, recipent, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address:    %s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipient_blockchain_address: %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value:                        %.1f\n", t.chanvalue)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.chanvalue,
	})
}
