package refund

type CreateRefundRequestPayload struct {
	RefID      string `json:"ref_id"`
	CustomerID string `json:"customer_id"`
	PaymentID  string `json:"payment_id"`
	Amount     string `json:"amount"`
	Notes      string `json:"notes"`
}
