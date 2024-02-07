package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()

	assert.Equal(t, len(privKey.Bytes()), privKeyLen)
	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), pubKeyLen)
}

func TestNewPrivateKeyFromString(t *testing.T) {
	var (
		seed       = "d6553d918962c38451e88f5ae7adb48b4fef8b766384e056255a91efa10470a4"
		privKey    = NewPrivateKeyFromString(seed)
		addressStr = "a46b03f373c5324a1e182ef5d3ee22b673155774"
	)
	// Generating a seed for the test
	// seed := make([]byte, 32)
	// io.ReadFull(rand.Reader, seed)
	// fmt.Println(hex.EncodeToString(seed))

	assert.Equal(t, privKeyLen, len(privKey.Bytes()))
	address := privKey.Public().Address()
	assert.Equal(t, addressStr, address.String())

}
func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	PubKey := privKey.Public()
	msg := []byte("foo bar baz")

	sig := privKey.Sign(msg)
	assert.True(t, sig.verify(PubKey, msg))

	//Test with invalid massage
	assert.False(t, sig.verify(PubKey, []byte("foo")))

	//Test with invalid pub key
	invalidPrivKey := GeneratePrivateKey()
	invalidPubKey := invalidPrivKey.Public()
	assert.False(t, sig.verify(invalidPubKey, msg))
}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubkey := privKey.Public()
	address := pubkey.Address()

	assert.Equal(t, addressLen, len(address.Bytes()))
	fmt.Println(address)
}
