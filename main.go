package main

import (
	"fmt"
	"log"
	"os"
	"persycoins/block"
	"persycoins/utils"
	"persycoins/wallet"
)

var logs utils.GoAppTools

func main() {
	//Criando o sistema de logs
	InfoLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)
	ErrorLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)

	logs.InfoLogger = *InfoLogger
	logs.ErrorLogger = *ErrorLogger

	fmt.Println(" ---- Teste Persy Coins -----")
	carteiraM := wallet.NewWallet()
	carteiraA := wallet.NewWallet()
	carteiraB := wallet.NewWallet()

	blockchain := block.NewBlockChain(carteiraM.Address)

	transaction := wallet.NewTransaction(carteiraA.PrivateKey, carteiraA.PublicKey, carteiraA.Address, carteiraB.Address, 1)
	blockchain.AddTransaction(transaction, carteiraA.PublicKey, logs)

	transaction = wallet.NewTransaction(carteiraB.PrivateKey, carteiraB.PublicKey, carteiraB.Address, carteiraA.Address, 2)
	blockchain.AddTransaction(transaction, carteiraB.PublicKey, logs)

	blockchain.Mining(carteiraM.PublicKey, carteiraM.PrivateKey, logs)
	blockchain.Print()

	fmt.Printf("A: %.1f\n", blockchain.CalculateTotalAmount(carteiraA.Address))
	fmt.Printf("B: %.1f\n", blockchain.CalculateTotalAmount(carteiraB.Address))
	fmt.Printf("M: %.1f\n", blockchain.CalculateTotalAmount(carteiraM.Address))
}
