package mdb

import "time"

// ProcessedChainTransaction is the durable idempotency claim for a chain
// transfer. It is separate from orders because one payment may be copied onto
// both a child order and its parent order after settlement.
type ProcessedChainTransaction struct {
	Network       string    `gorm:"column:network;size:32;uniqueIndex:processed_chain_tx_network_id_uindex,priority:1"`
	TransactionID string    `gorm:"column:transaction_id;size:256;uniqueIndex:processed_chain_tx_network_id_uindex,priority:2"`
	TradeID       string    `gorm:"column:trade_id;size:64;index:processed_chain_tx_trade_id_index"`
	CreatedAt     time.Time `gorm:"column:created_at"`
}

func (t *ProcessedChainTransaction) TableName() string {
	return "processed_chain_transactions"
}
