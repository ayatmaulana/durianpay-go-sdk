package disbursement

import (
	"context"

	"github.com/ayatmaulana/durianpay-go-sdk/common"
)

const (
  ROUTE_GROUP = "/disbursement"
)

type Disbursement struct {
  *common.Agent
}

func (d *Disbursement) FetchBankList(ctx context.Context) (res *ResponseBankList, err error) {
  err = d.Agent.Call(ctx, "GET", ROUTE_GROUP + "/banks", nil, res)
  return 
}

func (d *Disbursement) Topup(ctx context.Context, req *TopupRequestPayload) (res *TopupResponse, err error) {
  err = d.Agent.Call(ctx, "POST", ROUTE_GROUP + "/topup", req, res)
  return 
}

func (d *Disbursement) TopupDetailById(ctx context.Context, topupId string) (res *TopupHistoryResponse, err error) {
  err = d.Agent.Call(ctx, "GET", ROUTE_GROUP + "/topup/" + topupId, nil, res)
  return
}

func (d *Disbursement) FetchBalance(ctx context.Context) (res *FetchBalanceResponse, err error) {
  err = d.Agent.Call(ctx, "GET", ROUTE_GROUP + "/topup/balance", nil, res)
  return
}

func (d *Disbursement) SubmitDisbursement(ctx context.Context, req *SubmitDisbursementRequestPayload) (res *SubmitDisbursementResponse, err error) {
  err = d.Agent.Call(ctx, "POST", ROUTE_GROUP + "/submit", req, res)
  return
}

func (d *Disbursement) ApproveDisbursement(ctx context.Context, disbursementId string) (res *ApproveDisbursementResponse, err error){
  err = d.Agent.Call(ctx, "POST", ROUTE_GROUP + disbursementId + "/approve", nil, res)
  return
}

func (d *Disbursement) FetchDisbursementItemsById(ctx context.Context, disbursementId string) (res *FetchDisbursementItemsByIdResponse, err error) {
  err = d.Agent.Call(ctx, "GET", ROUTE_GROUP + disbursementId + "/items", nil, res)
  return
}

func (d *Disbursement) FetchDisbursementById(ctx context.Context, disbursementId string) (res *FetchDisbursementByIdResponse, err error) {
  err = d.Agent.Call(ctx, "GET", ROUTE_GROUP + disbursementId, nil, res)
  return
}
func (d *Disbursement) DeleteDisbursement(ctx context.Context, disbursementId string) (res *DeleteDisbursementResponse, err error){
  err = d.Agent.Call(ctx, "DELETE", ROUTE_GROUP + disbursementId, nil, res)
  return
}
func (d *Disbursement) ValidateBankAccount(ctx context.Context, req *ValidateBankACcountRequestPayload) (res *ValidateBankAccountResponse, err error) {
  err = d.Agent.Call(ctx, "POST", ROUTE_GROUP + "/validate", req, res)
  return
}

