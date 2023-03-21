package main

import (
	"fmt"
	"persycoins/block"
	"persycoins/wallet"
)

func main() {
	fmt.Println(" ---- Teste Persy Coins -----")
	carteiraM := wallet.NewWallet()
	carteiraA := wallet.NewWallet()
	carteiraB := wallet.NewWallet()

	transaction := wallet.NewTransaction(carteiraA.PrivateKey, carteiraA.PublicKey, carteiraA.Address, carteiraB.Address, 1)
	fmt.Printf("Essa é a assinatura da transação q ta no main: %s\n", transaction.GenerateSignature())
	blockchain := block.NewBlockChain(carteiraM.Address)
	isAdd := blockchain.AddTransaction(transaction, carteiraA.PublicKey)

	fmt.Println(isAdd)
}
