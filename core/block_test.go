package core

import (
	"github.com/ivnhk/blockchain-go/crypto"
	"github.com/ivnhk/blockchain-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.Hash{})
	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.Hash{})
	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}

func randomBlock(height uint32, prevBlockHash types.Hash) *Block {
	h := Header{
		Version:       1,
		DataHash:      types.RandomHash(),
		PrevBlockHash: prevBlockHash,
		Timestamp:     time.Now().UnixNano(),
		Height:        height,
	}
	return NewBlock(&h, []Transaction{})
}

func randomBlockWithSignature(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(height, prevBlockHash)
	tx := randomTxWithSignature(t)
	b.AddTransaction(tx)
	assert.Nil(t, b.Sign(privKey))
	return b
}
