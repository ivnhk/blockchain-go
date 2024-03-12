package core

type Blockchain struct {
	store     Storage
	headers   []*Header
	validator Validator
}

func (bc *Blockchain) AddBlock(b *Block) error {
	// validate that a block is valid

	//bc.store.append(b)
	//bc.Headers = append(bc.Headers, b.Header)

	return nil
}

func (bc *Blockchain) Height() uint32 {
	return uint32(len(bc.headers) - 1)
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{
		headers: []*Header{},
	}
	bc.validator = NewBlockValidator(bc)
	return bc
}
