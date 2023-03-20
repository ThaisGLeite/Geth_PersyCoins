package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
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

func NewBlockChain() *BlockChain {
	var newHash [32]byte
	blockchain := new(BlockChain)
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

func (blockchain *BlockChain) AddTransaction(sender string, receiver string, value float32) {
	transaction := NewTransaction(sender, receiver, value)
	blockchain.transactionPool = append(blockchain.transactionPool, transaction)
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

func main() {
	fmt.Println(" ---- Teste Persy Coins -----")
	blockchain := NewBlockChain()
	blockchain.Print()

	blockchain.AddTransaction("A", "B", 1.5)
	previousHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(5, previousHash)
	blockchain.Print()

	blockchain.AddTransaction("B", "C", 2)
	blockchain.AddTransaction("A", "C", 2)
	previousHash = blockchain.LastBlock().Hash()
	blockchain.CreateBlock(2, previousHash)
	blockchain.Print()
}
