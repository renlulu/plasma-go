package util

import "github.com/ethereum/go-ethereum/crypto/sha3"

func Sha3(seed string) [32]byte {
	return sha3.Sum256([]byte(seed))
}
