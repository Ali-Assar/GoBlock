package server

import (
	"context"

	"github.com/Ali-Assar/GoBlock/proto"
)

type Node struct {
	proto.UnimplementedNodeServer
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.None, error) {
	return nil, nil
}
