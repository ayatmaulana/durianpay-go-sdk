package settlement

import (
	"context"
	"net/http"

	"github.com/ayatmaulana/durianpay-go-sdk/common"
)

const (
  ROUTE_GROUP string = "/settlements"
)

type Settlement struct {
  *common.Agent
}

func (s *Settlement) FetchSettlement(ctx context.Context) (res *FetchSettlementResponse, err error) {
  err = s.Agent.Call(ctx, http.MethodGet, ROUTE_GROUP, nil, res)
  return
}

func (s *Settlement) FetchSettlementById(ctx context.Context, settlementId string) (res *FetchSettlementByIdResponse, err error) {
  err = s.Agent.Call(ctx, http.MethodGet, ROUTE_GROUP + "/" + settlementId , nil, res)
  return
}

func (s *Settlement) FetchSettlementDetail(ctx context.Context) (res *FetchSettlementDetailResponse, err error) {
  err = s.Agent.Call(ctx, http.MethodGet, ROUTE_GROUP + "/details", nil, res)
  return
}

func (s *Settlement) FetchSettlementDetailByPaymentId(ctx context.Context, paymentId string) (res *FetchSettlementDetailByPaymentIdResponse, err error) {
  err = s.Agent.Call(ctx, "GET", ROUTE_GROUP + "/details?payment_id=" + paymentId, nil, res)
  return
}

