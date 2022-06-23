package payment

import (
	"context"
	"net/http"

	"github.com/ayatmaulana/durianpay-go-sdk/common"
)

const (
  ROUTE_GROUP = "/payments"
)

type SandboxOption struct {
  ForceFail bool `json:"force_fail"`
  DelayMS int `json:"delay_ms"`
}

type Payment struct {
  *common.Agent
}


func (p *Payment) FetchPayment(ctx context.Context) (res *FetchPaymentResponse, err error) {
  res = &FetchPaymentResponse{}
  err = p.Agent.Call(ctx, http.MethodGet, ROUTE_GROUP, nil, res)
  return
}

func (p *Payment) FetchPaymentById(ctx context.Context, paymentId string) (res *FetchPaymentByIdResponse, err error) {
  err = p.Agent.Call(ctx, http.MethodGet, ROUTE_GROUP + paymentId, nil, res)
  return 
}


// Charge with any payment method
func (p *Payment) ChargePaymentVA(ctx context.Context, req *ChargePaymentVARequestPayload, option *SandboxOption) (res *ChargePaymentVAResponse, err error) {
  res = &ChargePaymentVAResponse{}
  payload := struct{
    Type string
    Request *ChargePaymentVARequestPayload
    SandboxOption *SandboxOption
  }{
    Type: "VA",
    Request: req,
  }

  if p.ClientConfig.Mode == common.SANDBOX {
    payload.SandboxOption = option
  }

  err = p.Agent.Call(ctx, http.MethodPost, ROUTE_GROUP + "/charge", payload, res)
  return
}

func (p *Payment) ChargePaymentEWallet(ctx context.Context, req *ChargePaymentEWalletRequestPayload) (res *ChargePaymentEWalletResponse, err error) {
  res = &ChargePaymentEWalletResponse{}
  payload := struct{
    Type string
    Request *ChargePaymentEWalletRequestPayload
  }{
    Type: "EWALLET",
    Request: req,
  }

  err = p.Agent.Call(ctx, http.MethodPost, ROUTE_GROUP + "/charge", payload, res)
  return
}

func (p *Payment) ChargePaymentRetailStore(ctx context.Context, req *ChargePaymentRetailStorePayload) (res *ChargePaymentRetailStorePayload, err error) {
  res = &ChargePaymentRetailStorePayload{}
  payload := struct{
    Type string
    Request *ChargePaymentRetailStorePayload
  }{
    Type: "RETAILSTORE",
    Request: req,
  }

  err = p.Agent.Call(ctx, http.MethodPost, ROUTE_GROUP + "/charge", payload, res)
  return
}

func (p *Payment) ChargePaymentOnlineCard(ctx context.Context, req *ChargePaymentCardRequestPayload) (res *ChargePaymentCardResponse, err error) {
  res = &ChargePaymentCardResponse{}
  payload := struct{
    Type string
    Request *ChargePaymentCardRequestPayload
  }{
    Type: "CARD",
    Request: req,
  }

  err = p.Agent.Call(ctx, http.MethodPost, ROUTE_GROUP + "/charge", payload, res)
  return
}

func (p *Payment) ChargePaymentOnlineBanking(ctx context.Context, req *ChargePaymentOnlineBankingRequestPayload) (res *ChargePaymentOnlineBankingRequestPayload, err error) {
  res = &ChargePaymentOnlineBankingRequestPayload{}
  payload := struct{
    Type string
    Request *ChargePaymentOnlineBankingRequestPayload
  }{
    Type: "ONLINE_BANKING",
    Request: req,
  }

  err = p.Agent.Call(ctx, http.MethodPost, ROUTE_GROUP + "/charge", payload, res)
  return
}

func (p *Payment) ChargePaymentQRIS(ctx context.Context, req *ChargePaymentQRISRequestPayload) (res *ChargePaymentQRISResponse, err error) {
  res = &ChargePaymentQRISResponse{}
  payload := struct{
    Type string
    Request *ChargePaymentQRISRequestPayload
  }{
    Type: "QRIS",
    Request: req,
  }
  err = p.Agent.Call(ctx, http.MethodPost, ROUTE_GROUP + "/charge", payload, res)
  return
}

func (p *Payment) ChargePaymentBNPL(ctx context.Context, req *ChargePaymentBNPLRequestPayload) (res *ChargePaymentBNPLResponse, err error) {
  res = &ChargePaymentBNPLResponse{}
  payload := struct{
    Type string
    Request *ChargePaymentBNPLRequestPayload
  }{
    Type: "BNPL",
    Request: req,
  }
  err = p.Agent.Call(ctx, http.MethodPost, ROUTE_GROUP + "/charge", payload, res)
  return
}

func (p *Payment) CheckPaymentStatus(ctx context.Context, paymentId string) (res *CheckPaymentStatusResponse, err error) {
  res = &CheckPaymentStatusResponse{}
  err = p.Agent.Call(ctx, http.MethodGet, ROUTE_GROUP + "/" + paymentId + "/status", nil, res)
  return
}

func (p *Payment) VerifyPayment(ctx context.Context, paymentId string) (res *VerifyPaymentResponse, err error)  {
  res = &VerifyPaymentResponse{}
  err = p.Agent.Call(ctx, http.MethodPost, ROUTE_GROUP + "/" + paymentId + "/verify", nil, res)
  return
}

func (p *Payment) CancelPayment(ctx context.Context, paymentId string) (res *CancelPaymentResponse, err error) {
  res = &CancelPaymentResponse{}
  err = p.Agent.Call(ctx, http.MethodPost, ROUTE_GROUP + "/" + paymentId + "/cancel", nil, res)
  return
}

func (p *Payment) MDRFeeCalculation(ctx context.Context, amount string, paymentMethod string) (res *MDRFeeCalculationResponse, err error){
  res = &MDRFeeCalculationResponse{}
  err = p.Agent.Call(ctx, http.MethodGet, "/merchants/mdr_fees", nil, res)
  return 
}
