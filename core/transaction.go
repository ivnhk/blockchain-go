package core

import (
	"github.com/ivnhk/blockchain-go/types"
	"io"
)

type Transaction struct {
	Data []byte
	From types.Address
}

func (tx *Transaction) EncodeBinary(w io.Writer) error {
	return nil
}

func (tx *Transaction) DecodeBinary(r io.Reader) error {
	return nil
}
