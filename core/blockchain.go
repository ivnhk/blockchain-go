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

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
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
