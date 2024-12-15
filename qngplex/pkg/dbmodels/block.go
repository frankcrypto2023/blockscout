package dbmodels

import (
	"time"

	"gorm.io/gorm"
)

const (
	BLOCK_STATUS_UNSTABLE = 0 // not stable
	BLOCK_STATUS_STABLE   = 1 // stable
	STABLE_THRESHOLD      = 10
)

func GetBlockStatus(corder, latestorder uint) uint {
	if corder > latestorder-STABLE_THRESHOLD {
		return BLOCK_STATUS_UNSTABLE
	}
	return BLOCK_STATUS_STABLE
}
func GetBlockColor(isblue uint) string {
	// 0:not blue;  1：blue  2：Cannot confirm
	switch isblue {
	case 0:
		return "red"
	case 1:
		return "blue"
	default:
		return "Cannot confirm"
	}
}

type UTXOBlock struct {
	gorm.Model
	BlockOrder        uint
	Version           uint
	Hash              string
	Height            uint
	Weight            uint
	Coinbase          uint // coinbase reward
	Confirmations     uint
	Nonce             uint64
	Pow               string
	BlockTime         time.Time
	Bits              uint64
	TransactionsCount uint
	ParentRoot        string
	StateRoot         string
	TxRoot            string
	Status            uint
	Txsvalid          uint
	Color             string
	Parents           string
	Children          string
}

func InitBlockTable(db *gorm.DB) {
	db.AutoMigrate(&UTXOBlock{})
}
