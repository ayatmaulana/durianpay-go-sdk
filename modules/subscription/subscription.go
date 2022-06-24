package subscription

import (
	"context"
	"net/http"

	"github.com/ayatmaulana/durianpay-go-sdk/common"
)

const ROUTE_GROUP = "/subscriptions"

type Subscription struct {
  *common.Agent
}

func (s *Subscription) CreateSubscriptionLink(ctx context.Context, req *CreateSubscriptionLinkRequestPayload) (res *CreateSubscriptionLinkResponse, err error) {
  res = &CreateSubscriptionLinkResponse{}
  err = s.Agent.Call(ctx, http.MethodPost, ROUTE_GROUP, req, res)
  return 
}

func (s *Subscription) FetchSubscription(ctx context.Context) (res *FetchSubscriptionResponse, err error) {
  res = &FetchSubscriptionResponse{}
  err = s.Agent.Call(ctx, http.MethodGet, ROUTE_GROUP, nil, res)
  return
}

func (s *Subscription) FetchSubscriptionById(ctx context.Context, subscriptionId string) (res *FetchSubscriptionByIdResponse, err error) {
  res = &FetchSubscriptionByIdResponse{}
  err = s.Agent.Call(ctx, http.MethodGet, ROUTE_GROUP + "/" + subscriptionId, nil, res)
  return
}

func (s *Subscription) CancelSubscription(ctx context.Context, subscriptionId string) (res *CancelSubscriptionResponse, err error) {
  res = &CancelSubscriptionResponse{}
  err = s.Agent.Call(ctx, http.MethodDelete, ROUTE_GROUP + "/" + subscriptionId + "/cancel", nil, res)
  return
}
