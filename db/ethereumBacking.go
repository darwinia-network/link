package db

import (
	"github.com/darwinia-network/link/util"
	"github.com/shopspring/decimal"
	"time"
)

type EthereumLockRecord struct {
	ID             uint            `gorm:"primary_key" json:"-"`
	CreatedAt      time.Time       `json:"-"`
	Chain          string          `json:"chain"`
	Tx             string          `json:"tx" sql:"size:100"`
	Token          string          `json:"token", sql:"size:100"`
	Sender         string          `json:"sender" sql:"size:100"`
	Receiver       string          `json:"receiver" sql:"size:100"`
	BlockNum       int             `json:"block_num"`
	Amount         decimal.Decimal `json:"amount" sql:"type:decimal(40,0);" `
	BlockTimestamp int             `json:"block_timestamp"`
	DarwiniaTx     string          `json:"darwinia_tx"`
	IsRelayed      bool            `json:"is_relayed" gorm:"-"`
}

func AddEthereumLockRecord(chain, tx, token, sender, receiver string, amount decimal.Decimal, blockNum, blockTimestamp int) error {
	db := util.DB
	query := db.Create(&EthereumLockRecord{
        Chain: chain, Tx: tx, Token: token, Sender: sender, Receiver: receiver, Amount: amount, BlockNum: blockNum,
		BlockTimestamp: blockTimestamp,
	})
	return query.Error
}

func EthereumLockList(sender string, page, row int) (list []EthereumLockRecord, count int) {
	util.DB.Model(EthereumLockRecord{}).Where("sender = ?", sender).Count(&count)
	util.DB.Where("sender = ?", sender).Order("block_num desc").Offset(page * row).Limit(row).Find(&list)

	for index, value := range list {
		list[index].IsRelayed = GetRelayBestBlockNum() >= uint64(value.BlockNum)
	}
	return list, count
}

func UpdateEthereumLockRecord(tx, darwiniaTx string) {
	util.DB.Model(EthereumLockRecord{}).Where("tx = ?", tx).Update(EthereumLockRecord{DarwiniaTx: darwiniaTx})
}

