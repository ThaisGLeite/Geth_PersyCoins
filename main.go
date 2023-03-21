package main

import (
	"fmt"
	"persycoins/wallet"
)

func main() {
	fmt.Println(" ---- Teste Persy Coins -----")
	carteira := wallet.NewWallet()
	fmt.Println(carteira.PrivateKey)
	fmt.Println(carteira.PublicKey)
	fmt.Println(carteira.Address)

	transaction := wallet.NewTransaction(carteira.PrivateKey, carteira.PublicKey, carteira.Address, "B", 1.5)
	fmt.Printf("assinatura %s\n", transaction.GenerateSignature())

}
