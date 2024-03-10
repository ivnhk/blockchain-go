package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"github.com/ivnhk/blockchain-go/types"
	"io"
)

var (
	ByteOrder = binary.LittleEndian
)

type Header struct {
	Version   uint32
	PrevBlock types.Hash
	Timestamp int64
	Height    uint32
	Nonce     uint64
}

func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, ByteOrder, &h.Version); err != nil {
		return err
	}
	if err := binary.Write(w, ByteOrder, &h.PrevBlock); err != nil {
		return err
	}
	if err := binary.Write(w, ByteOrder, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Write(w, ByteOrder, &h.Height); err != nil {
		return err
	}
	return binary.Write(w, ByteOrder, &h.Nonce)
}

func (h *Header) DecodeBinary(r io.Reader) error {
	if err := binary.Read(r, ByteOrder, &h.Version); err != nil {
		return err
	}
	if err := binary.Read(r, ByteOrder, &h.PrevBlock); err != nil {
		return err
	}
	if err := binary.Read(r, ByteOrder, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Read(r, ByteOrder, &h.Height); err != nil {
		return err
	}
	return binary.Read(r, ByteOrder, &h.Nonce)
}

type Block struct {
	Header
	Transactions []Transaction
	// cached version of the header hash
	hash types.Hash
}

func (b Block) Hash() types.Hash {
	if b.hash.IsZero() {
		buf := &bytes.Buffer{}
		b.Header.EncodeBinary(buf)
		b.hash = types.Hash(sha256.Sum256(buf.Bytes()))
	}
	
	return b.hash
}

func (b *Block) EncodeBinary(w io.Writer) error {
	if err := b.Header.EncodeBinary(w); err != nil {
		return err
	}

	for _, tx := range b.Transactions {
		if err := tx.EncodeBinary(w); err != nil {
			return err
		}
	}
	return nil
}

func (b *Block) DecodeBinary(r io.Reader) error {
	if err := b.Header.DecodeBinary(r); err != nil {
		return err
	}

	for _, tx := range b.Transactions {
		if err := tx.DecodeBinary(r); err != nil {
			return err
		}
	}
	return nil

}
