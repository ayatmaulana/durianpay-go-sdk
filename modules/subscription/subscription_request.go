package subscription

import "time"

type CreateSubscriptionLinkRequestPayload struct {
	Plan              Plan      `json:"plan"`
	Customer          Customer  `json:"customer"`
	BillingCycleCount int       `json:"billing_cycle_count"`
	StartedAt         time.Time `json:"started_at"`
	ChargeType        string    `json:"charge_type"`
	Notes             string    `json:"notes"`
}
type Plan struct {
	SubscriptionRefID  string `json:"subscription_ref_id"`
	Name               string `json:"name"`
	BillingPeriod      string `json:"billing_period"`
	BillingPeriodCount int    `json:"billing_period_count"`
	GracePeriod        string `json:"grace_period"`
	GracePeriodCount   int    `json:"grace_period_count"`
	Price              string `json:"price"`
	Qty                string `json:"qty"`
}
type Customer struct {
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	GivenName  string `json:"given_name"`
	MiddleName string `json:"middle_name"`
	SurName    string `json:"sur_name"`
}
