package types

import (
	"testing"

	"github.com/Ali-Assar/GoBlock/util"
	"github.com/stretchr/testify/assert"
)

func TestHashBlok(t *testing.T) {
	block := util.RandomBlock()
	hash := HashBlock(block)

	assert.Equal(t, 32, len(hash))
}
