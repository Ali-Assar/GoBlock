package main

import (
	"context"
	"log"

	"github.com/Ali-Assar/GoBlock/node"
	"github.com/Ali-Assar/GoBlock/proto"
	"google.golang.org/grpc"
)

func main() {
	makeNode(":3000", []string{})
	makeNode(":4000", []string{":3000"})

	// go func() {
	// 	for {
	// 		time.Sleep(2 * time.Second)
	// 		makeTransaction()
	// 	}
	// }()
	select {}
}

func makeNode(listenAddr string, bootstrapNodes []string) *node.Node {
	n := node.NewNode()
	go n.Start(listenAddr)
	if len(bootstrapNodes) > 0 {
		if err := n.BootStrapNetwork(bootstrapNodes); err != nil {
			log.Fatal(err)
		}
	}
	return n
}

func makeTransaction() {
	client, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewNodeClient(client)

	version := &proto.Version{
		Version:    "goBlock-0.1",
		Height:     1,
		ListenAddr: ":4000",
	}
	_, err = c.Handshake(context.TODO(), version)
	if err != nil {
		log.Fatal(err)
	}
}
