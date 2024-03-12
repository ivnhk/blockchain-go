package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockchain(t *testing.T) {
	bc := NewBlockchain()
	assert.NotNil(t, bc.validator)
}
