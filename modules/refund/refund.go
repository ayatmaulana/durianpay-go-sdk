package refund

import (
	"context"
	"net/http"

	"github.com/ayatmaulana/durianpay-go-sdk/common"
)

const (
  ROUTE_GROUP string = "/refunds"
)

type Refund struct {
  *common.Agent
}

func (r *Refund) CreateRefund(ctx context.Context, req *CreateRefundRequestPayload) (res *CreateRefundResponse, err error) {
  res = &CreateRefundResponse{}
  err = r.Agent.Call(ctx, http.MethodPost, ROUTE_GROUP, req, res)
  return
}

func (r *Refund) FetchRefund(ctx context.Context) (res *FetchRefundResponse, err error) {
  res = &FetchRefundResponse{}
  err = r.Agent.Call(ctx, http.MethodGet, ROUTE_GROUP, nil, res)
  return
}

func (r *Refund) FetchRefundById(ctx context.Context, refundId string) (res *FetchRefundByIdResponse, err error) {
  res = &FetchRefundByIdResponse{}
  err = r.Agent.Call(ctx, http.MethodGet, ROUTE_GROUP + "/" + refundId, nil, res)
  return
}
