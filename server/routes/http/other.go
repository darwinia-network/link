package http

import (
	"github.com/darwinia-network/link/services/contract"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"math/big"
	"net/http"
)

func supply() gin.HandlerFunc {
	return func(c *gin.Context) {
		supply := new(big.Int).Add(contract.RingEthSupply(), contract.RingTronSupply())
		c.JSON(http.StatusOK, JsonFormat(decimal.NewFromBigInt(supply, 0).Div(decimal.New(1, 18)), 0))
	}
}
