package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"goblockchain/utils"
)

type Transaction struct {
	txid                       string
	senderPrivateKey           *ecdsa.PrivateKey
	senderPublicKey            *ecdsa.PublicKey
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey,
	sender string, recipient string, value float32) *Transaction {
	txID := createTxID(sender, recipient, value)
	return &Transaction{txID, privateKey, publicKey, sender, recipient, value}
}

func (t *Transaction) GetTxID() *string {
	return &t.txid
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

func (t *Transaction) GenerateSignature() *utils.Signature {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	r, s, _ := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
	return &utils.Signature{R: r, S: s}
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TxID      string  `json:"txid"`
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		TxID:      t.txid,
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}
