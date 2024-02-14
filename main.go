package main

import (
	"context"
	"log"
	"time"

	"github.com/Ali-Assar/GoBlock/crypto"
	"github.com/Ali-Assar/GoBlock/node"
	"github.com/Ali-Assar/GoBlock/proto"
	"github.com/Ali-Assar/GoBlock/util"
	"google.golang.org/grpc"
)

func main() {
	makeNode(":3000", []string{}, true)
	time.Sleep(time.Second)
	makeNode(":4000", []string{":3000"}, false)
	time.Sleep(time.Second)
	makeNode(":5000", []string{":4000"}, false)

	for {
		time.Sleep(2 * time.Second)
		makeTransaction()
	}
}

func makeNode(listenAddr string, bootstrapNodes []string, isValidator bool) *node.Node {
	cfg := node.ServerConfig{
		Version:    "GoBlock-1",
		ListenAddr: listenAddr,
	}
	if isValidator {
		cfg.PrivateKey = crypto.GeneratePrivateKey()
	}
	n := node.NewNode(cfg)
	go n.Start(listenAddr, bootstrapNodes)
	return n
}

func makeTransaction() {
	client, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewNodeClient(client)
	privKey := crypto.GeneratePrivateKey()
	tx := &proto.Transaction{
		Version: 1,
		Inputs: []*proto.TxInput{
			{
				PrevTxHash:   util.RandomHash(),
				PrevOutIndex: 0,
				PublicKey:    privKey.Public().Address().Bytes(),
			},
		},
		Outputs: []*proto.TxOutput{
			{
				Amount:  99,
				Address: privKey.Public().Bytes(),
			},
		},
	}
	_, err = c.HandleTransaction(context.TODO(), tx)
	if err != nil {
		log.Fatal(err)
	}
}
