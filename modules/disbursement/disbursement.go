package disbursement

import (
	"context"

	"github.com/ayatmaulana/durianpay-go-sdk/common"
)

const (
  ROUTE_GROUP = "/disbursement"
)

type Disburesement struct {
  *common.Agent
}

func (d *Disburesement) FetchBankList(ctx context.Context) (res *ResponseBankList, err error) {
  d.Agent.Call(ctx, "GET",  ROUTE_GROUP + "/banks", nil)
  return 
}


func Topup() {

}

func TopupDetailById() {

}

func FetchBalance() {

}

func SubmitDisbursement() {

}

func ApproveDisbursement() {

}

func FetchDisburesementItemsById() {

}
