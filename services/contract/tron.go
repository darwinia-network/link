package contract

import (
	"github.com/darwinia-network/link/lib/web3"
	"github.com/darwinia-network/link/util"
	"math/big"
)

var tronContract = "416e0d26adf5323f5b82d5714354dc3c6870adee7c"

type TronResponse struct {
	ConstantResult []string `json:"constant_result"`
}

func RingTronSupply() *big.Int {
	w := web3.New("tron")
	var e TronResponse
	_ = w.Call(&e, tronContract, "totalSupply()")
	return util.U256(e.ConstantResult[0])
}
