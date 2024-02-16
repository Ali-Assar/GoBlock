package types

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/Ali-Assar/GoBlock/crypto"
	"github.com/Ali-Assar/GoBlock/proto"
	"github.com/Ali-Assar/GoBlock/util"
	"github.com/stretchr/testify/assert"
)

func TestCalculateRootHash(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	block := util.RandomBlock()
	tx := &proto.Transaction{
		Version: 1,
	}
	block.Transactions = append(block.Transactions, tx)
	SignBlock(privKey, block)

	assert.True(t, VerifyRootHash(block))
	assert.Equal(t, 32, len(block.Header.RootHash))

}

func TestSignVerifyBlock(t *testing.T) {
	var (
		block   = util.RandomBlock()
		privKey = crypto.GeneratePrivateKey()
		pubKey  = privKey.Public()

		sig = SignBlock(privKey, block)
	)

	assert.Equal(t, 64, len(sig.Bytes()))
	assert.True(t, sig.Verify(pubKey, HashBlock(block)))

	assert.Equal(t, block.PublicKey, pubKey.Bytes())
	assert.Equal(t, block.Signature, sig.Bytes())
	assert.True(t, VerifyBlock(block))

	invalidPrivKey := crypto.GeneratePrivateKey()
	block.PublicKey = invalidPrivKey.Public().Bytes()
	assert.False(t, VerifyBlock(block))
}
func TestHashBlok(t *testing.T) {
	block := util.RandomBlock()
	hash := HashBlock(block)
	fmt.Println(hex.EncodeToString(hash))

	assert.Equal(t, 32, len(hash))
}
