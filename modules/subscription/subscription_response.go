package subscription

import "time"

type CreateSubscriptionLinkResponse struct {
	Data struct {
		ID   string `json:"id"`
		Plan struct {
			ID                 string `json:"id"`
			Name               string `json:"name"`
			BillingPeriod      string `json:"billing_period"`
			BillingPeriodCount int    `json:"billing_period_count"`
			GracePeriod        string `json:"grace_period"`
			GracePeriodCount   int    `json:"grace_period_count"`
			Price              string `json:"price"`
			Qty                string `json:"qty"`
			Amount             string `json:"amount"`
			SubscriptionRefID  string `json:"subscription_ref_id"`
		} `json:"plan"`
		Customer struct {
			ID         string `json:"id"`
			Email      string `json:"email"`
			Mobile     string `json:"mobile"`
			GivenName  string `json:"given_name"`
			MiddleName string `json:"middle_name"`
			SurName    string `json:"sur_name"`
		} `json:"customer"`
		Status             string    `json:"status"`
		StartedAt          time.Time `json:"started_at"`
		NextDueAt          time.Time `json:"next_due_at"`
		SubscriptionOrders []struct {
			BillingPeriod int    `json:"billing_period"`
			OrderID       string `json:"order_id"`
			Status        string `json:"status"`
		} `json:"subscription_orders"`
		ChargeType string `json:"charge_type"`
		Link       string `json:"link"`
	} `json:"data"`
}

type FetchSubscriptionResponse struct {
	Data struct {
		Subscription []struct {
			ID                 string    `json:"id"`
			Status             string    `json:"status"`
			ChargeType         string    `json:"charge_type"`
			Name               string    `json:"name"`
			StartedAt          time.Time `json:"started_at"`
			BillingPeriod      string    `json:"billing_period"`
			BillingPeriodCount int       `json:"billing_period_count"`
			Customer           struct {
				ID         string `json:"id"`
				Email      string `json:"email"`
				Mobile     string `json:"mobile"`
				GivenName  string `json:"given_name"`
				MiddleName string `json:"middle_name"`
				SurName    string `json:"sur_name"`
			} `json:"customer"`
		} `json:"subscription"`
		Count int `json:"count"`
	} `json:"data"`
}

type FetchSubscriptionByIdResponse struct {
	Data struct {
		ID   string `json:"id"`
		Plan struct {
			ID                 string `json:"id"`
			Name               string `json:"name"`
			BillingPeriod      string `json:"billing_period"`
			BillingPeriodCount int    `json:"billing_period_count"`
			GracePeriod        string `json:"grace_period"`
			GracePeriodCount   int    `json:"grace_period_count"`
			Price              string `json:"price"`
			SubscriptionRefID  string `json:"subscription_ref_id"`
		} `json:"plan"`
		ChargeType        string    `json:"charge_type"`
		Status            string    `json:"status"`
		StartedAt         time.Time `json:"started_at"`
		NextDueDate       time.Time `json:"next_due_date"`
		BillingCycleCount int       `json:"billing_cycle_count"`
		Amount            string    `json:"amount"`
		ActiveOrderID     string    `json:"active_order_id"`
		Link              string    `json:"link"`
	} `json:"data"`
}

type CancelSubscriptionResponse struct {
	Data struct {
		ID          string    `json:"id"`
		Status      string    `json:"status"`
		StartedAt   time.Time `json:"started_at"`
		CancelledAt time.Time `json:"cancelled_at"`
	} `json:"data"`
}
