package disbursement

import (
	"context"
	"net/http"

	"github.com/ayatmaulana/durianpay-go-sdk/common"
)

const (
  ROUTE_GROUP = "/disbursements"
)

type Disbursement struct {
  *common.Agent
}

func (d *Disbursement) FetchBankList(ctx context.Context) (res *ResponseBankList, err error) {
  res = &ResponseBankList{}
  err = d.Agent.Call(ctx, "GET", ROUTE_GROUP + "/banks", nil, &res)
  return
}

func (d *Disbursement) Topup(ctx context.Context, req *TopupRequestPayload) (res *TopupResponse, err error) {
  res = &TopupResponse{}
  err = d.Agent.Call(ctx, "POST", ROUTE_GROUP + "/topup", req, res)
  return 
}

func (d *Disbursement) TopupDetailById(ctx context.Context, topupId string) (res *TopupHistoryResponse, err error) {
  res = &TopupHistoryResponse{}
  err = d.Agent.Call(ctx, "GET", ROUTE_GROUP + "/topup/" + topupId, nil, res)
  return
}

func (d *Disbursement) FetchBalance(ctx context.Context) (res *FetchBalanceResponse, err error) {
  res = &FetchBalanceResponse{}
  err = d.Agent.Call(ctx, "GET", ROUTE_GROUP + "/topup/balance", nil, res)
  return
}

func (d *Disbursement) SubmitDisbursement(ctx context.Context, req *SubmitDisbursementRequestPayload) (res *SubmitDisbursementResponse, err error) {
  res = &SubmitDisbursementResponse{}
  err = d.Agent.Call(ctx, "POST", ROUTE_GROUP + "/submit", req, res)
  return
}

func (d *Disbursement) ApproveDisbursement(ctx context.Context, disbursementId string) (res *ApproveDisbursementResponse, err error){
  res = &ApproveDisbursementResponse{}
  err = d.Agent.Call(ctx, http.MethodPost, ROUTE_GROUP + disbursementId + "/approve", nil, res)
  return
}

func (d *Disbursement) FetchDisbursementItemsById(ctx context.Context, disbursementId string) (res *FetchDisbursementItemsByIdResponse, err error) {
  res = &FetchDisbursementItemsByIdResponse{}
  err = d.Agent.Call(ctx, http.MethodGet, ROUTE_GROUP + disbursementId + "/items", nil, res)
  return
}

func (d *Disbursement) FetchDisbursementById(ctx context.Context, disbursementId string) (res *FetchDisbursementByIdResponse, err error) {
  res = &FetchDisbursementByIdResponse{}
  err = d.Agent.Call(ctx, http.MethodGet, ROUTE_GROUP + disbursementId, nil, res)
  return
}

func (d *Disbursement) DeleteDisbursement(ctx context.Context, disbursementId string) (res *DeleteDisbursementResponse, err error){
  res = &DeleteDisbursementResponse{}
  err = d.Agent.Call(ctx, http.MethodDelete, ROUTE_GROUP + disbursementId, nil, res)
  return
}

func (d *Disbursement) ValidateBankAccount(ctx context.Context, req *ValidateBankACcountRequestPayload) (res *ValidateBankAccountResponse, err error) {
  res = &ValidateBankAccountResponse{}
  err = d.Agent.Call(ctx, http.MethodPost, ROUTE_GROUP + "/validate", req, res)
  return
}

