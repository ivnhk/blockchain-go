package core

import (
	"github.com/ivnhk/blockchain-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
		block := randomBlockWithSignature(t, uint32(i), getPrevBlockHash(t, bc, uint32(i)))
		assert.Nil(t, bc.AddBlock(block))
	}
	assert.Equal(t, uint32(lenBlocks), bc.Height())
	assert.Equal(t, lenBlocks+1, len(bc.headers))
	assert.NotNil(t, bc.AddBlock(randomBlock(89, types.Hash{})))
}

func TestAddBlockTooHigh(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, 3, types.Hash{})))
}

func TestGetHeader(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	lenBlocks := 1000
	for i := 1; i <= lenBlocks; i++ {
		block := randomBlockWithSignature(t, uint32(i), getPrevBlockHash(t, bc, uint32(i)))
		assert.Nil(t, bc.AddBlock(block))
		header, err := bc.GetHeader(block.Height)
		assert.Nil(t, err)
		assert.Equal(t, block.Header, header)
	}
}

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0, types.Hash{}))
	assert.Nil(t, err)

	return bc
}

func getPrevBlockHash(t *testing.T, bc *Blockchain, height uint32) types.Hash {
	prevBlockHeader, err := bc.GetHeader(height - 1)
	assert.Nil(t, err)

	return BlockHasher{}.Hash(prevBlockHeader)
}
