package core

import (
	"fmt"
	"log/slog"
)

type Blockchain struct {
	store     Storage
	headers   []*Header
	validator Validator
}

func (bc *Blockchain) AddBlock(b *Block) error {
	if err := bc.validator.ValidateBlock(b); err != nil {
		return err
	}

	bc.addBlockWithoutValidation(b)
	return nil
}

func (bc *Blockchain) Height() uint32 {
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

func (bc *Blockchain) GetHeader(height uint32) (*Header, error) {
	if height > bc.Height() {
		return nil, fmt.Errorf("given height (%d) too high", height)
	}
	return bc.headers[height], nil
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	slog.Info(
		"Add block without validation",
		"height", b.Height,
		"hash", b.Hash(BlockHasher{}),
	)
	bc.headers = append(bc.headers, b.Header)

	return bc.store.Put(b)
}

func NewBlockchain(genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		headers: []*Header{},
		store:   NewMemoryStore(),
	}
	bc.validator = NewBlockValidator(bc)

	err := bc.addBlockWithoutValidation(genesis)
	return bc, err
}
