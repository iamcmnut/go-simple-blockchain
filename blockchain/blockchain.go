package blockchain

import "fmt"

type Blockchain struct {
	Blocks []*Block
}

func (chain *Blockchain) Init() {
	geneisBlock := &Block{
		Data:              "this is genesisblock",
		Nonce:             0,
		PreviousBlockHash: "0000000000000000000000000000000000000000000000000000000000000000",
	}

	geneisBlock.GetValidNonce()

	chain.Blocks = []*Block{geneisBlock}
}

func (chain *Blockchain) ValidateChain() bool {
	for _, b := range chain.Blocks {
		if !b.Validate() {
			return false
		}
	}

	return true
}

func (chain *Blockchain) GetLatestBlock() (*Block, error) {
	l := len(chain.Blocks)

	if l > 0 {
		return chain.Blocks[l-1], nil
	}

	return nil, fmt.Errorf("Empty chain")
}

func (chain *Blockchain) GetBlock(index int) (*Block, error) {
	l := len(chain.Blocks)

	if index < l && index >= 0 {
		return chain.Blocks[index], nil
	}

	return nil, fmt.Errorf("Block not found")
}

func (chain *Blockchain) CreateBlock(data string) (*Block, error) {
	latestBlock, err := chain.GetLatestBlock()
	if err != nil {
		return nil, err
	}

	newBlock := &Block{
		Data:              data,
		PreviousBlockHash: latestBlock.GetHash(),
		Nonce:             0,
	}

	newBlock.GetValidNonce()
	chain.Blocks = append(chain.Blocks, newBlock)
	return newBlock, nil
}
