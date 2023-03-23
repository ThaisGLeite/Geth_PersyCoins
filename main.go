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

	blockchain := block.NewBlockChain(carteiraM.Address)

	transaction := wallet.NewTransaction(carteiraA.PrivateKey, carteiraA.PublicKey, carteiraA.Address, carteiraB.Address, 1)
	isAdd := blockchain.AddTransaction(transaction, carteiraA.PublicKey)

	fmt.Println(isAdd)

	transaction = wallet.NewTransaction(carteiraB.PrivateKey, carteiraB.PublicKey, carteiraB.Address, carteiraA.Address, 2)
	isAdd = blockchain.AddTransaction(transaction, carteiraB.PublicKey)
	fmt.Println(isAdd)
}
