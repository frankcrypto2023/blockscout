package dbmodels

import (
	"time"

	"gorm.io/gorm"
)

const (
	SpentStatusUnspent  = 0
	SpentStatusSpent    = 1
	SpentStatusImmature = 2
)

type UTXOTx struct {
	gorm.Model
	Txid        string
	Block       string
	BlockOrder  uint
	IsCoinbase  uint
	PreTxs      string
	Amount      uint
	Confirm     uint
	TxTime      time.Time
	TxSize      uint
	Fee         uint
	Txvalid     uint
	SpentStatus uint // 0 unspent 1 spent 2 immature
	PKScript    string
	Vout        uint
	Address     string
}

func GetMatureStatus(mature uint, corder, latestorder uint) uint {
	if corder < latestorder-mature {
		return SpentStatusImmature
	}
	return SpentStatusUnspent
}
func InitTxTable(db *gorm.DB) {
	db.AutoMigrate(&UTXOTx{})
}
