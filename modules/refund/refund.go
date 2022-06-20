package refund

import (
	"context"

	"github.com/ayatmaulana/durianpay-go-sdk/common"
)

const (
  ROUTE_GROUP string = "/refunds"
)

type Refund struct {
  *common.Agent
}

func (r *Refund) CreateRefund(ctx context.Context, req *CreateRefundRequestPayload) (res *CreateRefundResponse, err error) {
  err = r.Agent.Call(ctx, "POST", ROUTE_GROUP, req, res)
  return
}

func (r *Refund) FetchRefund(ctx context.Context) (res *FetchRefundResponse, err error) {
  err = r.Agent.Call(ctx, "GET", ROUTE_GROUP, nil, res)
  return
}

func (r *Refund) FetchRefundById(ctx context.Context, refundId string) (res *FetchRefundByIdResponse, err error) {
  err = r.Agent.Call(ctx, "GET", ROUTE_GROUP + "/" + refundId, nil, res)
  return
}
