package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Nonce        int      `json:"nonce"`
	PreviousHash [32]byte `json:"previous_hash"`
	Timestamp    int64    `json:"timestamp"`
	Transactions []string `json:"transactions"`
}

type BlockChain struct {
	transactionPool []string
	chain           []*Block
}

func (blockchain *BlockChain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	bloco := NewBlock(nonce, previousHash)
	blockchain.chain = append(blockchain.chain, bloco)
	return bloco
}

func NewBlockChain() *BlockChain {
	var newHash [32]byte
	blockchain := new(BlockChain)
	blockchain.CreateBlock(0, newHash)
	return blockchain
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

func (b *Block) Print() {
	fmt.Printf("Timestamp          %d\n", b.Timestamp)
	fmt.Printf("Nonce              %d\n", b.Nonce)
	fmt.Printf("Previous Hash      %x\n", b.PreviousHash)
	fmt.Printf("Transactions       %s\n", b.Transactions)

}

func (bloco *Block) Hash() [32]byte {
	jsonB, _ := json.Marshal(bloco)
	return sha256.Sum256([]byte(jsonB))
}

func NewBlock(nonce int, previousHash [32]byte) *Block {
	b := new(Block)
	b.Timestamp = time.Now().UnixNano()
	b.Nonce = nonce
	b.PreviousHash = previousHash
	return b
}

func main() {
	fmt.Println(" ---- Teste Persy Coins -----")
	blockchain := NewBlockChain()
	blockchain.Print()

	previousHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(5, previousHash)
	blockchain.Print()

	previousHash = blockchain.LastBlock().Hash()
	blockchain.CreateBlock(2, previousHash)
	blockchain.Print()
}
