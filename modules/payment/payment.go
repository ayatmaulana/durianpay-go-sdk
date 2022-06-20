package payment

import (
	"context"

	"github.com/ayatmaulana/durianpay-go-sdk/common"
)

const (
  ROUTE_GROUP = "/payments"
)


type Payment struct {
  *common.Agent
}


func (p *Payment) FetchPayment(ctx context.Context) (res *FetchPaymentResponse, err error) {
  err = p.Agent.Call(ctx, "GET", ROUTE_GROUP, nil, res)
  return
}

func (p *Payment) FetchPaymentById(ctx context.Context, paymentId string) (res *FetchPaymentByIdResponse, err error) {
  err = p.Agent.Call(ctx, "GET", ROUTE_GROUP + paymentId, nil, res)
  return 
}


// Charge with any payment method
func (p *Payment) ChargePaymentVA(ctx context.Context, req *ChargePaymentVARequestPayload) (res *ChargePaymentVAResponse, err error) {
  payload := struct{
    Type string
    Request *ChargePaymentVARequestPayload
  }{
    Type: "VA",
    Request: req,
  }

  err = p.Agent.Call(ctx, "POST", ROUTE_GROUP + "/charge", payload, res)
  return
}

func (p *Payment) ChargePaymentEWallet(ctx context.Context, req *ChargePaymentEWalletRequestPayload) (res *ChargePaymentEWalletResponse, err error) {
  payload := struct{
    Type string
    Request *ChargePaymentEWalletRequestPayload
  }{
    Type: "EWALLET",
    Request: req,
  }

  err = p.Agent.Call(ctx, "POST", ROUTE_GROUP + "/charge", payload, res)
  return
}

func (p *Payment) ChargePaymentRetailStore(ctx context.Context, req *ChargePaymentRetailStorePayload) (res *ChargePaymentRetailStorePayload, err error) {
  payload := struct{
    Type string
    Request *ChargePaymentRetailStorePayload
  }{
    Type: "RETAILSTORE",
    Request: req,
  }

  err = p.Agent.Call(ctx, "POST", ROUTE_GROUP + "/charge", payload, res)
  return
}

func (p *Payment) ChargePaymentOnlineCard(ctx context.Context, req *ChargePaymentCardRequestPayload) (res *ChargePaymentCardResponse, err error) {
  payload := struct{
    Type string
    Request *ChargePaymentCardRequestPayload
  }{
    Type: "CARD",
    Request: req,
  }

  err = p.Agent.Call(ctx, "POST", ROUTE_GROUP + "/charge", payload, res)
  return
}

func (p *Payment) ChargePaymentOnlineBanking(ctx context.Context, req *ChargePaymentOnlineBankingRequestPayload) (res *ChargePaymentOnlineBankingRequestPayload, err error) {
  payload := struct{
    Type string
    Request *ChargePaymentOnlineBankingRequestPayload
  }{
    Type: "ONLINE_BANKING",
    Request: req,
  }

  err = p.Agent.Call(ctx, "POST", ROUTE_GROUP + "/charge", payload, res)
  return
}

func (p *Payment) ChargePaymentQRIS(ctx context.Context, req *ChargePaymentQRISRequestPayload) (res *ChargePaymentQRISResponse, err error) {
  payload := struct{
    Type string
    Request *ChargePaymentQRISRequestPayload
  }{
    Type: "QRIS",
    Request: req,
  }
  err = p.Agent.Call(ctx, "POST", ROUTE_GROUP + "/charge", payload, res)
  return
}

func (p *Payment) ChargePaymentBNPL(ctx context.Context, req *ChargePaymentBNPLRequestPayload) (res *ChargePaymentBNPLResponse, err error) {
  payload := struct{
    Type string
    Request *ChargePaymentBNPLRequestPayload
  }{
    Type: "BNPL",
    Request: req,
  }
  err = p.Agent.Call(ctx, "POST", ROUTE_GROUP + "/charge", payload, res)
  return
}



func (p *Payment) CheckPaymentStatus(ctx context.Context, paymentId string) (res *CheckPaymentStatusResponse, err error) {
  err = p.Agent.Call(ctx, "GET", ROUTE_GROUP + "/" + paymentId + "/status", nil, res)
  return
}

func (p *Payment) VerifyPayment(ctx context.Context, paymentId string) (res *VerifyPaymentResponse, err error)  {
  err = p.Agent.Call(ctx, "POST", ROUTE_GROUP + "/" + paymentId + "/verify", nil, res)
  return
}

func (p *Payment) CancelPayment(ctx context.Context, paymentId string) (res *CancelPaymentResponse, err error) {
  err = p.Agent.Call(ctx, "POST", ROUTE_GROUP + "/" + paymentId + "/cancel", nil, res)
  return
}

func (p *Payment) MDRFeeCalculation(ctx context.Context, amount string, paymentMethod string) (res *MDRFeeCalculationResponse, err error){
  err = p.Agent.Call(ctx, "GET", "/merchants/mdr_fees", nil, res)
  return 
}
