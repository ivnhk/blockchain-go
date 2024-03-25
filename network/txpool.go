package network

import (
	"github.com/ivnhk/blockchain-go/core"
	"github.com/ivnhk/blockchain-go/types"
)

type TxPool struct {
	transactions map[types.Hash]*core.Transaction
}

func (p *TxPool) Len() int {
	return len(p.transactions)
}

func (p *TxPool) Flush() {
	p.transactions = make(map[types.Hash]*core.Transaction)
}

// Add adds a transaction to the pool. The caller is responsible for checking if transaction already exists
func (p *TxPool) Add(tx *core.Transaction) error {
	hash := tx.Hash(core.TxHasher{})
	p.transactions[hash] = tx
	return nil
}

func (p *TxPool) Has(hash types.Hash) bool {
	_, ok := p.transactions[hash]
	return ok
}

func NewTxPool() *TxPool {
	return &TxPool{
		transactions: make(map[types.Hash]*core.Transaction),
	}
}
