package order

import (
	"context"
	"net/http"

	"github.com/ayatmaulana/durianpay-go-sdk/common"
)

const (
  ROUTE_GROUP = "/orders"
)


type Order struct {
  *common.Agent
}

func (o *Order) CreateOrder(ctx context.Context, req *CreateOrderRequestPayload) (res *CreateOrderResponse, err error) {
  res = &CreateOrderResponse{}
  o.Agent.Call(ctx, http.MethodPost, ROUTE_GROUP, req, res)
  return
}

func (o *Order) FetchOrder(ctx context.Context) (res *FetchOrderResponse, err error) {
  res = &FetchOrderResponse{}
  o.Agent.Call(ctx, http.MethodGet, ROUTE_GROUP, nil, res)
  return
}

func (o *Order) FetchOrderById(ctx context.Context, orderId string) (res *FetchOrderByIdResponse, err error) {
  res = &FetchOrderByIdResponse{}
  o.Agent.Call(ctx, http.MethodGet, ROUTE_GROUP + orderId, nil, res)
  return
}

func (o *Order) CreateInstapayOrPaymentLink(ctx context.Context, req *CreateInstapayOrPaymentLinkRequestPayload) (res *CreateInstapayOrPaymentLinkResponse, err error) {
  res = &CreateInstapayOrPaymentLinkResponse{}
  o.Agent.Call(ctx, http.MethodPost, ROUTE_GROUP, req, res)
  return
}
