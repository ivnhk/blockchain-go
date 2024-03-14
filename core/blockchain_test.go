package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockchain(t *testing.T) {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)
	assert.NotNil(t, bc.validator)
}
