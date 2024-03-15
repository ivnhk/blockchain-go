package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)

	return bc
}

func TestNewBlockchain(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	lenBlocks := 1000
	for i := 1; i <= lenBlocks; i++ {
		block := randomBlock(uint32(i))
		assert.Nil(t, bc.AddBlock(block))
	}
	assert.Equal(t, uint32(lenBlocks), bc.Height())
	assert.Equal(t, lenBlocks+1, len(bc.headers))
	assert.NotNil(t, bc.AddBlock(randomBlock(89)))
}
