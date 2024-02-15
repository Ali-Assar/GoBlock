package node

import (
	"testing"

	"github.com/Ali-Assar/GoBlock/types"
	"github.com/Ali-Assar/GoBlock/util"
	"github.com/stretchr/testify/assert"
)

func TestNewChain(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore())
	assert.Equal(t, 0, chain.Height())
	_, err := chain.GetBlockByHeight(0)
	assert.Nil(t, err)
}

func TestChainHeight(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore())
	for i := 0; i < 100; i++ {
		b := util.RandomBlock()

		prevBlock, err := chain.GetBlockByHeight(chain.Height())
		assert.Nil(t, err)
		b.Header.PrevHash = types.HashHeader(prevBlock.Header)

		assert.Nil(t, chain.AddBlock(b))
		assert.Equal(t, chain.Height(), i+1)
	}
}

func TestAddBlock(t *testing.T) {
	chain := NewChain(NewMemoryBlockStore())

	for i := 0; i < 100; i++ {

		block := util.RandomBlock()
		prevBlock, err := chain.GetBlockByHeight(chain.Height())
		assert.Nil(t, err)
		block.Header.PrevHash = types.HashHeader(prevBlock.Header)
		blockHash := types.HashBlock(block)

		assert.Nil(t, chain.AddBlock(block))

		fetchedBlock, err := chain.GetBlockByHash(blockHash)
		assert.Nil(t, err)
		assert.Equal(t, block, fetchedBlock)

		fetchedBlockByHeight, err := chain.GetBlockByHeight(i + 1)
		assert.Nil(t, err)
		assert.Equal(t, block, fetchedBlockByHeight)
	}
}
