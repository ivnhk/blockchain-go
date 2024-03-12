package core

import (
	"github.com/ivnhk/blockchain-go/crypto"
	"github.com/ivnhk/blockchain-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func randomBlock(height uint32) *Block {
	h := Header{
		Version:       1,
		DataHash:      types.RandomHash(),
		PrevBlockHash: types.Hash{},
		Timestamp:     time.Now().UnixNano(),
		Height:        height,
	}
	tx := Transaction{
		Data: []byte("lol"),
	}
	return NewBlock(&h, []Transaction{tx})
}

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)
	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)
	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}
