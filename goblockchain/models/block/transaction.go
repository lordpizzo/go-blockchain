package block

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	txid                       string
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(txID string, sender string, recipient string, value float32) *Transaction {
	return &Transaction{txID, sender, recipient, value}
}

func createTxID(sender string, recipient string, value float32) string {
	// Concatenar os dados da transação em uma única string
	stringToHash := sender + recipient + fmt.Sprintf("%.1f", value)

	// Criar um hash SHA-256 da string concatenada
	hash := sha256.New()
	hash.Write([]byte(stringToHash))
	txid := hash.Sum(nil)

	// Converter o hash para uma string hexadecimal
	return hex.EncodeToString(txid)
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address      %s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipient_blockchain_address   %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value                          %.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Txid      string  `json:"txid"`
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Txid:      t.txid,
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	v := &struct {
		Txid      *string  `json:"txid"`
		Sender    *string  `json:"sender_blockchain_address"`
		Recipient *string  `json:"recipient_blockchain_address"`
		Value     *float32 `json:"value"`
	}{
		Txid:      &t.txid,
		Sender:    &t.senderBlockchainAddress,
		Recipient: &t.recipientBlockchainAddress,
		Value:     &t.value,
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	return nil
}
