package order

import (
	"context"

	"github.com/ayatmaulana/durianpay-go-sdk/common"
)

const (
  ROUTE_GROUP = "/orders"
)


type Order struct {
  *common.Agent
}

func (o *Order) CreateOrder(ctx context.Context, req *CreateOrderRequestPayload) (res *CreateOrderResponse, err error) {
  o.Agent.Call(ctx, "POST", ROUTE_GROUP, req, res)
  return
}

func (o *Order) FetchOrder(ctx context.Context) (res *FetchOrderResponse, err error) {
  o.Agent.Call(ctx, "GET", ROUTE_GROUP, nil, res)
  return
}

func (o *Order) FetchOrderById(ctx context.Context, orderId string) (res *FetchOrderByIdResponse, err error) {
  o.Agent.Call(ctx, "GET", ROUTE_GROUP + orderId, nil, res)
  return
}

func (o *Order) CreateInstapayOrPaymentLink(ctx context.Context, req *CreateInstapayOrPaymentLinkRequestPayload) (res *CreateInstapayOrPaymentLinkResponse, err error) {
  o.Agent.Call(ctx, "POST", ROUTE_GROUP, req, res)
  return
}
