package node

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/Ali-Assar/GoBlock/crypto"
	"github.com/Ali-Assar/GoBlock/proto"
	"github.com/Ali-Assar/GoBlock/types"
)

type HeaderList struct {
	headers []*proto.Header
}

func NewHeaderList() *HeaderList {
	return &HeaderList{
		headers: []*proto.Header{},
	}
}

func (list *HeaderList) Add(h *proto.Header) {
	list.headers = append(list.headers, h)
}

func (list *HeaderList) Get(index int) *proto.Header {
	if index > list.Height() {
		panic("index too high!")
	}
	return list.headers[index]
}

func (list *HeaderList) Height() int {
	return list.Len() - 1
}

func (list *HeaderList) Len() int {
	return len(list.headers)
}

type Chain struct {
	BlockStore BlockStorer
	headers    *HeaderList
}

func NewChain(bs BlockStorer) *Chain {
	chain := &Chain{
		BlockStore: bs,
		headers:    NewHeaderList(),
	}
	chain.addBlock(createGenesisBlock())
	return chain
}

func (c *Chain) Height() int {
	return c.headers.Height()
}

// add bloc without validation
func (c *Chain) AddBlock(b *proto.Block) error {
	if err := c.ValidateBlock(b); err != nil {
		return err
	}
	return c.addBlock(b)
}

// add bloc with validation
func (c *Chain) addBlock(b *proto.Block) error {
	// Add the header to the list of headers
	c.headers.Add(b.Header)
	// validation
	return c.BlockStore.Put(b)
}

func (c *Chain) GetBlockByHash(hash []byte) (*proto.Block, error) {
	hashHex := hex.EncodeToString(hash)
	return c.BlockStore.Get(hashHex)
}

func (c *Chain) GetBlockByHeight(height int) (*proto.Block, error) {
	if c.Height() < height {
		return nil, fmt.Errorf("given height (%d) too high - height (%d)", height, c.Height())
	}
	header := c.headers.Get(height)
	hash := types.HashHeader(header)
	return c.GetBlockByHash(hash)
}

func (c *Chain) ValidateBlock(b *proto.Block) error {
	currentBlock, err := c.GetBlockByHeight(c.Height())
	if err != nil {
		return err
	}
	hash := types.HashBlock(currentBlock)
	if !bytes.Equal(hash, b.Header.PrevHash) {
		return fmt.Errorf("invalid previous block hash")
	}

	return nil
}

func createGenesisBlock() *proto.Block {
	privKey := crypto.GeneratePrivateKey()
	block := &proto.Block{
		Header: &proto.Header{
			Version: 1,
		},
	}
	types.SignBlock(privKey, block)
	return block
}
