package node

import (
	"testing"

	"github.com/Ali-Assar/GoBlock/types"
	"github.com/Ali-Assar/GoBlock/util"
	"github.com/stretchr/testify/assert"
)

func TestAddBlock(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore())
	block := util.RandomBlock()
	blockHash := types.HashBlock(block)

	assert.Nil(t, chain.AddBlock(block))
	fetchedBlock, err := chain.GetBlockByHash(blockHash)
	assert.Nil(t, err)
	assert.Equal(t, block, fetchedBlock)
}
