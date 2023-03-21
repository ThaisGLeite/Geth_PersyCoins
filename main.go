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

	transaction := wallet.NewTransaction(carteiraA.PrivateKey, carteiraA.PublicKey, carteiraA.Address, carteiraB.Address, 1.5)

	blockchain := block.NewBlockChain(carteiraM.Address)
	isAdd := blockchain.AddTransaction(carteiraA.Address, carteiraB.Address, 2, carteiraA.PublicKey, transaction.GenerateSignature())

	fmt.Println(isAdd)
}
