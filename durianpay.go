package durianpay

import (
  "github.com/ayatmaulana/durianpay-go-sdk/common"

  "github.com/ayatmaulana/durianpay-go-sdk/modules/disbursement"
  "github.com/ayatmaulana/durianpay-go-sdk/modules/order"
)

type Modules struct {
  Disbursement disbursement.Disburesement
  Order order.Order
}

func NewClient(config *common.ClientConfig) *Modules {
  agent := common.NewAgent(config)

  return &Modules{
    Disbursement: disbursement.Disburesement{Agent: agent},
    Order: order.Order{Agent: agent},
  }
}
