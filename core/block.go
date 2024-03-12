package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/ivnhk/blockchain-go/crypto"
	"github.com/ivnhk/blockchain-go/types"
	"io"
)

type Header struct {
	Version       uint32
	DataHash      types.Hash
	PrevBlockHash types.Hash
	Timestamp     int64
	Height        uint32
}

type Block struct {
	*Header
	Transactions []Transaction
	Validator    crypto.PublicKey
	Signature    *crypto.Signature
	// cached version of the header hash
	hash types.Hash
}

func (b *Block) Encode(r io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(r, b)
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Hash(hr Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hr.Hash(b)
	}
	return b.hash
}

func (b *Block) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(b.HeaderData())
	if err != nil {
		return err
	}
	b.Validator = privKey.PublicKey()
	b.Signature = sig
	return nil
}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block has no signature")
	}
	if !b.Signature.Verify(b.Validator, b.HeaderData()) {
		return fmt.Errorf("invalid block signature")
	}
	return nil
}

func (b *Block) HeaderData() []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	enc.Encode(b.Header)

	return buf.Bytes()
}

func NewBlock(h *Header, txx []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: txx,
	}
}
