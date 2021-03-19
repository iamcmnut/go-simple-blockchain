package blockchain

import (
	"crypto/sha256"
	"fmt"
	"math"
)

type Block struct {
	Data              string `json:"data"`
	PreviousBlockHash string `json:"previousBlockHash"`
	Nonce             uint32 `json:"nonce"`
}

func (b *Block) GetHash() string {
	data := fmt.Sprintf("%010d%s%s", b.Nonce, b.PreviousBlockHash, b.Data)
	h := sha256.New()
	h.Write([]byte(data))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return hash
}

func (b *Block) GetValidNonce(prevNonce uint32) bool {
	var nonce float64 = 1
	var pNonce float64 = float64(prevNonce)
	for b.GetHash()[:4] != "0000" {
		newNonce := math.Pow(nonce, 2) - math.Pow(pNonce, 2)
		b.Nonce = uint32(newNonce)
		nonce += 1
	}

	return true
}

func (b *Block) Validate() bool {
	if b.GetHash()[:4] != "0000" {
		return false
	}

	return true
}
