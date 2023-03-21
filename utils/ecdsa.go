package utils

import "math/big"

type Signature struct {
	R *big.Int `json:"receiver_addres"`
	S *big.Int `json:"sender_addres"`
}
