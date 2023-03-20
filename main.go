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
}
