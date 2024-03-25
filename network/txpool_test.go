package network

import (
	"github.com/ivnhk/blockchain-go/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	p := NewTxPool()
	assert.Equal(t, 0, p.Len())
}

func TestAddTx(t *testing.T) {
	p := NewTxPool()
	tx := core.NewTransaction([]byte("lol"))
	assert.Nil(t, p.Add(tx))
	assert.Equal(t, 1, p.Len())

	p.Flush()
	assert.Equal(t, 0, p.Len())
}
