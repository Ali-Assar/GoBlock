package node

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/Ali-Assar/GoBlock/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type Node struct {
	version    string
	listenAddr string

	peerLock sync.RWMutex
	peers    map[proto.NodeClient]*proto.Version

	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	return &Node{
		peers:   make(map[proto.NodeClient]*proto.Version),
		version: "goBlock-0.1",
	}
}

func (n *Node) addPeer(c proto.NodeClient, v *proto.Version) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()

	fmt.Printf("new peer connected (%s) - height(%d)\n", v.ListenAddr, v.Height)
	n.peers[c] = v
}

func (n *Node) removePeer(c proto.NodeClient) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()
	delete(n.peers, c)
}

func (n *Node) bootStrapNetwork(addrs []string) error {
	for _, addr := range addrs {
		c, err := makeNodeClient(addr)
		if err != nil {
			return err
		}

		v, err := c.Handshake(context.Background(), n.getVersion())
		if err != nil {
			fmt.Println("handshake error:", err)
			continue
		}

		n.addPeer(c, v)
	}
	return nil
}

func (n *Node) Start(listenAddr string) error {
	n.listenAddr = listenAddr
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	proto.RegisterNodeServer(grpcServer, n)
	fmt.Println("node running on port ", listenAddr)
	return grpcServer.Serve(ln)
}

func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {
	fromVersion := &proto.Version{
		Version: n.version,
		Height:  100,
	}

	c, err := makeNodeClient(v.ListenAddr)
	if err != nil {
		return nil, err
	}

	n.addPeer(c, v)

	return fromVersion, nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println("received tx from:", peer)
	return &proto.Ack{}, nil
}

func (n *Node) getVersion() *proto.Version {
	return &proto.Version{
		Version:    "goBlock-0.1",
		Height:     0,
		ListenAddr: n.listenAddr,
	}
}

func makeNodeClient(listenAddr string) (proto.NodeClient, error) {
	conn, err := grpc.Dial(listenAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return proto.NewNodeClient(conn), nil
}
