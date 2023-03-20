package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"math/big"

	"github.com/btcsuite/btcutil/base58"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    string
}

type Transaction struct {
	SenderPrivateKey *ecdsa.PrivateKey
	SenderPublicKey  *ecdsa.PublicKey
	SenderAdress     string
	ReceiverAddress  string
	Value            float32
}

type Signature struct {
	R *big.Int `json:"receiver_addres"`
	S *big.Int `json:"sender_addres"`
}

// São 8 passos para criar a carteira, suas chaves e o seu endereço e o seu endereço
func NewWallet() *Wallet {
	// 1 - Criar uuma chave privada ECDSA de 32 bytes e uma chave publica de 64 bytes
	carteira := new(Wallet)
	carteira.PrivateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	carteira.PublicKey = &carteira.PrivateKey.PublicKey

	// 2 - Fazer o hash SHA256 da chave publica (32 bytes)
	hash := sha256.New()
	hash.Write(carteira.PublicKey.X.Bytes())
	hash.Write(carteira.PublicKey.Y.Bytes())
	passo2 := hash.Sum(nil)

	// 3 - Executar a função de criptografia RIPEMD-160 e fazer o hash do resultado
	hash = sha256.New()
	hash.Write(passo2)
	passo3 := hash.Sum(nil)

	// 4 - Adicionar os bytes 0x00 na frente do hash do passo 3
	passo4 := make([]byte, 21)
	passo4[0] = 0x00
	copy(passo4[1:], passo3)

	// 5 - Fazer o hash no resultado do passo 4
	hash = sha256.New()
	hash.Write(passo4)
	passo5 := hash.Sum(nil)

	// 6 - Fazer hash do resultado do passo 4, ou seja, o hash do hash
	hash = sha256.New()
	hash.Write(passo5)
	passo6 := hash.Sum(nil)

	// 7 - Pegar os primeiros 4 bytes do passo 6 para realizar a checksum (verificação da soma dos bits)
	checksum := passo6[:4]
	passo7 := make([]byte, 25)
	copy(passo7[:21], passo4)
	copy(passo7[21:], checksum)

	// 8 - Converter o resultado para uma string de byte q use a base58 e guardar o endereço na carteira
	carteira.Address = base58.Encode(passo7)

	return carteira
}

func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, sender string, receiver string, value float32) *Transaction {
	return &Transaction{
		SenderPrivateKey: privateKey,
		SenderPublicKey:  publicKey,
		SenderAdress:     sender,
		ReceiverAddress:  receiver,
		Value:            value,
	}
}

func (transaction *Transaction) GenerateSignature() *Signature {
	json, _ := json.Marshal(transaction)
	hash := sha256.Sum256([]byte(json))
	r, s, _ := ecdsa.Sign(rand.Reader, transaction.SenderPrivateKey, hash[:])
	//Monta a estrutura da assinatura para retornar
	return &Signature{r, s}
}
