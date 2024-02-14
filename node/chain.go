package node

import (
	"encoding/hex"

	"github.com/Ali-Assar/GoBlock/proto"
)

type HeaderList struct {
	headers []*proto.Header
}

func (list *HeaderList) Add(h *proto.Header) {
	list.headers = append(list.headers, h)
}

func (list *HeaderList) Height() int {
	return list.Len() - 1
}

func (list *HeaderList) Len() int {
	return len(list.headers)
}

type Chain struct {
	BlockStore BlockStorer
}

func NewChain(bs BlockStorer) *Chain {
	return &Chain{
		BlockStore: bs,
	}
}

func (c *Chain) AddBlock(b *proto.Block) error {
	return c.BlockStore.Put(b)
}

func (c *Chain) GetBlockByHash(hash []byte) (*proto.Block, error) {
	hashHex := hex.EncodeToString(hash)
	return c.BlockStore.Get(hashHex)
}

func (c *Chain) GetBlockByHeight(height int) (*proto.Block, error) {
	return nil, nil
}
