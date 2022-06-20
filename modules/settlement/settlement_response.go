package settlement

import "time"

type FetchSettlementResponse struct {
	TotalCount       int `json:"total_count"`
	SettlementDetail []struct {
		ID                     string    `json:"id"`
		SettlementAmount       string    `json:"settlement_amount"`
		Status                 string    `json:"status"`
		Fee                    string    `json:"fee"`
		PromoAmount            string    `json:"promo_amount"`
		TotalTranscationAmount string    `json:"total_transcation_amount"`
		CreatedAt              time.Time `json:"created_at"`
		SettledAt              time.Time `json:"settled_at"`
		Currency               string    `json:"currency"`
	} `json:"settlement_detail"`
}

type FetchSettlementByIdResponse struct {
	Data struct {
		ID               string    `json:"id"`
		SettlementAmount string    `json:"settlement_amount"`
		Status           string    `json:"status"`
		Fee              string    `json:"fee"`
		CreatedAt        time.Time `json:"created_at"`
		SettledAt        time.Time `json:"settled_at"`
		PromoAmount      string    `json:"promo_amount"`
	} `json:"data"`
}

type FetchSettlementDetailResponse struct {
	SettlementDetail []struct {
		SettlementID       string    `json:"settlement_id"`
		PaymentID          string    `json:"payment_id"`
		PaymentReference   string    `json:"payment_reference"`
		OrderID            string    `json:"order_id"`
		OrderReference     string    `json:"order_reference"`
		Status             string    `json:"status"`
		Currency           string    `json:"currency"`
		SettlementAmount   string    `json:"settlement_amount"`
		TotalSettlementFee string    `json:"total_settlement_fee"`
		PaymentDiscount    string    `json:"payment_discount"`
		SettledAt          time.Time `json:"settled_at"`
		Group              string    `json:"group"`
		PaymentAmount      string    `json:"payment_amount"`
		PaymentDate        time.Time `json:"payment_date"`
		TransactionAmount  string    `json:"transaction_amount"`
		PaymentDetailsType string    `json:"payment_details_type"`
		PaymentMethodID    string    `json:"payment_method_id"`
	} `json:"settlement_detail"`
	SettlementCount  int `json:"settlement_count"`
	TransactionCount int `json:"transaction_count"`
}

type FetchSettlementDetailByPaymentIdResponse struct {
	Data struct {
		SettlementID       string    `json:"settlement_id"`
		PaymentID          string    `json:"payment_id"`
		PaymentReference   string    `json:"payment_reference"`
		OrderID            string    `json:"order_id"`
		OrderReference     string    `json:"order_reference"`
		Status             string    `json:"status"`
		Currency           string    `json:"currency"`
		SettlementAmount   string    `json:"settlement_amount"`
		TotalSettlementFee string    `json:"total_settlement_fee"`
		PaymentDiscount    string    `json:"payment_discount"`
		SettledAt          time.Time `json:"settled_at"`
		Group              string    `json:"group"`
		PaymentAmount      string    `json:"payment_amount"`
		PaymentDate        time.Time `json:"payment_date"`
		TransactionAmount  string    `json:"transaction_amount"`
		PaymentChannel     string    `json:"payment_channel"`
		PaymentSubchannel  string    `json:"payment_subchannel"`
	} `json:"data"`
}
