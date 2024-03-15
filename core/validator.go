package core

import "fmt"

type Validator interface {
	ValidateBlock(*Block) error
}

type BlockValidator struct {
	bc *Blockchain
}

func (v *BlockValidator) ValidateBlock(b *Block) error {
	if v.bc.HasBlock(b.Height) {
		return fmt.Errorf("chain already contains block (%d)", b.Height)
	}

	if b.Height != v.bc.Height()+1 {
		return fmt.Errorf("block (%d) too high", b.Height)
	}

	if err := b.Verify(); err != nil {
		return err
	}
	return nil
}

func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{
		bc: bc,
	}
}
