package blockchain

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Data              string
	PreviousBlockHash string
	Nonce             uint32
}

func (b *Block) GetHash() string {
	data := fmt.Sprintf("%010d%s%s", b.Nonce, b.PreviousBlockHash, b.Data)
	h := sha256.New()
	h.Write([]byte(data))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return hash
}

func (b *Block) GetValidNonce() bool {
	for b.GetHash()[:4] != "0000" {
		b.Nonce += 1
	}

	return true
}

func (b *Block) Validate() bool {
	if b.GetHash()[:4] != "0000" {
		return false
	}

	return true
}
