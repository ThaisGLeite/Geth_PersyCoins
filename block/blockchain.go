package block

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"persycoins/utils"
	"persycoins/wallet"
	"strings"
	"time"
)

const (
	MINING_DIFFICULTY = 3
	MINING_REWARD     = 1
	MINING_SENDER     = "Persy"
)

type Block struct {
	Nonce        int             `json:"nonce"`
	PreviousHash [32]byte        `json:"previous_hash"`
	Timestamp    int64           `json:"timestamp"`
	Transactions []*Transactions `json:"transactions"`
}

type BlockChain struct {
	transactionPool []*Transactions
	chain           []*Block
	addres          string
}

type Transactions struct {
	Sender   string  `json:"sender_address"`
	Receiver string  `json:"receiver_address"`
	Value    float32 `json:"value"`
}

func (blockchain *BlockChain) CreateBlock(nonce int, previousHash [32]byte) *Block {

	//Botar o pool de transações dentro do novo bloco
	bloco := NewBlock(nonce, previousHash, blockchain.transactionPool)
	blockchain.chain = append(blockchain.chain, bloco)

	//Apagar o pool de transações porque colocou eleas no bloco novo
	blockchain.transactionPool = []*Transactions{}
	return bloco
}

func NewBlockChain(address string) *BlockChain {
	var newHash [32]byte
	blockchain := new(BlockChain)
	blockchain.addres = address
	blockchain.CreateBlock(0, newHash)
	return blockchain
}

func NewTransaction(sender string, receiver string, value float32) *Transactions {
	transaction := Transactions{
		Sender:   sender,
		Receiver: receiver,
		Value:    value,
	}
	return &transaction
}

func (blockchain *BlockChain) LastBlock() *Block {
	ultimo := len(blockchain.chain) - 1
	return blockchain.chain[ultimo]
}

func (blockchain *BlockChain) Print() {
	for i, bloco := range blockchain.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		bloco.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 60))
}

func (blockchain *BlockChain) VerifyTransaction(senderKey *ecdsa.PublicKey, transaction *wallet.Transaction, logs utils.GoAppTools) bool {
	transactionJ, err := json.Marshal(transaction)
	utils.Check(err, logs)
	transactionH := sha256.Sum256([]byte(transactionJ))
	assinatura := transaction.GenerateSignature(logs)
	// Verificar a chave para retornar o resultado bool se a assinatura é verdadeira
	return ecdsa.Verify(senderKey, transactionH[:], assinatura.R, assinatura.S)
}

func (blockchain *BlockChain) AddTransaction(transaction *wallet.Transaction, senderKey *ecdsa.PublicKey, logs utils.GoAppTools) bool {

	if transaction.SenderAdress == MINING_SENDER {
		tempTransaction := &Transactions{
			Sender:   transaction.SenderAdress,
			Receiver: transaction.ReceiverAddress,
			Value:    transaction.Value,
		}
		blockchain.transactionPool = append(blockchain.transactionPool, tempTransaction)
		return true
	}

	if blockchain.VerifyTransaction(senderKey, transaction, logs) {
		tempTransaction := &Transactions{
			Sender:   transaction.SenderAdress,
			Receiver: transaction.ReceiverAddress,
			Value:    transaction.Value,
		}
		blockchain.transactionPool = append(blockchain.transactionPool, tempTransaction)
		return true
	} else {
		log.Println("Erro ao verificar a assinatura da transação")
		return false
	}
}

func (blockchain *BlockChain) CopyTransactionPool() []*Transactions {
	transactions := make([]*Transactions, 0)
	for _, t := range blockchain.transactionPool {
		novaTransacao := NewTransaction(t.Sender, t.Receiver, t.Value)
		transactions = append(transactions, novaTransacao)
	}
	return transactions
}

func (blockchain *BlockChain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transactions) bool {
	zeros := strings.Repeat("0", MINING_DIFFICULTY)
	guess := NewBlock(nonce, previousHash, transactions)
	guessHash := fmt.Sprintf("%x", guess.Hash())
	if guessHash[:MINING_DIFFICULTY] == zeros {
		return true
	} else {
		return false
	}
}

func (blockchain *BlockChain) ProofOfWork() int {
	transactions := blockchain.CopyTransactionPool()
	previousHash := blockchain.LastBlock().Hash()
	nonce := 0

	//Vai ficar tentando adivinhar o numero novo até conseguir
	for !(blockchain.ValidProof(nonce, previousHash, transactions)) {
		nonce++
	}
	return nonce
}

func (blockchain *BlockChain) Mining(minerPublicKey *ecdsa.PublicKey, minerPrivateKey *ecdsa.PrivateKey, logs utils.GoAppTools) bool {

	//TODO arrumar esse Mining pq ele num ta servido de nada, mas o add transaction
	//  teve q ser mudado do original pq a chave validava errado
	tempTransaction := &wallet.Transaction{
		SenderPublicKey:  minerPublicKey,
		SenderPrivateKey: minerPrivateKey,
		Value:            1,
	}
	blockchain.AddTransaction(tempTransaction, minerPublicKey, logs)
	nonce := blockchain.ProofOfWork()
	previousHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(nonce, previousHash)
	fmt.Println("Mining Success!")
	return true
}

func (blockchain *BlockChain) CalculateTotalAmount(address string) float32 {
	var total float32
	for _, bloco := range blockchain.chain {
		for _, transaction := range bloco.Transactions {
			if address == transaction.Receiver {
				total += transaction.Value
			}
			if address == transaction.Sender {
				total -= transaction.Value
			}
		}
	}
	return total
}

func (bloco *Block) Print() {
	fmt.Printf("Timestamp          %d\n", bloco.Timestamp)
	fmt.Printf("Nonce              %d\n", bloco.Nonce)
	fmt.Printf("Previous Hash      %x\n", bloco.PreviousHash)
	for _, transaction := range bloco.Transactions {
		transaction.Print()
	}
}

func (transaction *Transactions) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("Sender BlockChain Address:      %s\n", transaction.Sender)
	fmt.Printf("Receiver BlockChain Address:    %s\n", transaction.Receiver)
	fmt.Printf("Value:                          %.2f\n", transaction.Value)
}

func (bloco *Block) Hash() [32]byte {
	jsonB, _ := json.Marshal(bloco)
	return sha256.Sum256([]byte(jsonB))
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transactions) *Block {
	b := new(Block)
	b.Timestamp = time.Now().UnixNano()
	b.Nonce = nonce
	b.PreviousHash = previousHash
	b.Transactions = transactions
	return b
}
