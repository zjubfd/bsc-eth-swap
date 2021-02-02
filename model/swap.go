package model

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/binance-chain/bsc-eth-swap/common"
)

type SwapStartTxLog struct {
	Id    int64
	Chain string `gorm:"not null;index:swap_start_tx_log_chain"`

	ContractAddress string `gorm:"not null"`
	FromAddress     string `gorm:"not null"`
	ToAddress       string `gorm:"not null"`
	Amount          string `gorm:"not null"`
	FeeAmount       string `gorm:"not null"`

	Status       TxStatus `gorm:"not null;index:swap_start_tx_log_status"`
	TxHash       string   `gorm:"not null;index:swap_start_tx_log_tx_hash"`
	BlockHash    string   `gorm:"not null"`
	Height       int64    `gorm:"not null"`
	ConfirmedNum int64    `gorm:"not null"`

	Phase TxPhase `gorm:"not null;index:swap_start_tx_log_phase"`

	UpdateTime int64
	CreateTime int64
}

func (SwapStartTxLog) TableName() string {
	return "swap_start_txs"
}

func (l *SwapStartTxLog) BeforeCreate() (err error) {
	l.CreateTime = time.Now().Unix()
	l.UpdateTime = time.Now().Unix()
	return nil
}

type SwapFillTx struct {
	gorm.Model

	Direction         common.SwapDirection `gorm:"not null"`
	StartSwapTxHash   string               `gorm:"not null;index:swap_fill_tx_start_swap_tx_hash"`
	FillSwapTxHash    string               `gorm:"not null;index:swap_fill_tx_fill_swap_tx_hash"`
	GasPrice          string               `gorm:"not null"`
	ConsumedFeeAmount string
	Height            int64
	Status            FillTxStatus `gorm:"not null"`
	TrackRetryCounter int64
}

func (SwapFillTx) TableName() string {
	return "swap_fill_txs"
}

type Swap struct {
	gorm.Model

	Status common.SwapStatus `gorm:"not null;index:swap_status"`
	// the user addreess who start this swap
	Sponsor string `gorm:"not null;index:swap_sponsor"`

	BscContractAddr string `gorm:"not null;index:swap_bsc_contract_addr"`
	EThContractAddr string `gorm:"not null;index:swap_eth_contract_addr"`
	Symbol          string
	Amount          string               `gorm:"not null;index:swap_amount"`
	Decimals        int                  `gorm:"not null"`
	Direction       common.SwapDirection `gorm:"not null"`

	// The tx hash confirmed deposit
	StartTxHash string `gorm:"not null;index:swap_start_tx_hash"`
	// The tx hash confirmed withdraw
	FillTxHash string `gorm:"not null;index:swap_fill_tx_hash"`

	// The tx hash of refund
	RefundTxHash string

	// used to log more message about how this swap failed or invalid
	Log string

	RecordHash string `gorm:"not null"`
}

func (Swap) TableName() string {
	return "swaps"
}