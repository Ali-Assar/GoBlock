package types

import (
	"crypto/sha256"

	"github.com/Ali-Assar/GoBlock/crypto"
	"github.com/Ali-Assar/GoBlock/proto"
	pb "google.golang.org/protobuf/proto"
)

func SignBlock(pk *crypto.PrivateKey, b *proto.Block) *crypto.Signature {
	hash := HashBlock(b)
	sig := pk.Sign(hash)
	b.PublicKey = pk.Public().Bytes()
	b.Signature = sig.Bytes()
	return sig
}

// HashBlock returns a SHA256 of the header
func HashBlock(block *proto.Block) []byte {
	return HashHeader(block.Header)
}

func HashHeader(header *proto.Header) []byte {
	b, err := pb.Marshal(header)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(b)

	return hash[:]
}
