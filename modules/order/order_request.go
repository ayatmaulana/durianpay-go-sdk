package order

import "time"
type CreateOrderRequestPayload struct {
	Amount        string    `json:"amount"`
	PaymentOption string    `json:"payment_option"`
	Currency      string    `json:"currency"`
	OrderRefID    string    `json:"order_ref_id"`
	Customer      Customer  `json:"customer"`
	Items         []Items   `json:"items"`
	Metadata      Metadata  `json:"metadata"`
	ExpiryDate    time.Time `json:"expiry_date"`
}
type Address struct {
	ReceiverName  string `json:"receiver_name"`
	ReceiverPhone string `json:"receiver_phone"`
	Label         string `json:"label"`
	AddressLine1  string `json:"address_line_1"`
	AddressLine2  string `json:"address_line_2"`
	City          string `json:"city"`
	Region        string `json:"region"`
	Country       string `json:"country"`
	PostalCode    string `json:"postal_code"`
	Landmark      string `json:"landmark"`
}
type Customer struct {
	CustomerRefID string  `json:"customer_ref_id"`
	GivenName     string  `json:"given_name"`
	Email         string  `json:"email"`
	Mobile        string  `json:"mobile"`
	Address       Address `json:"address"`
}
type Items struct {
	Name  string `json:"name"`
	Qty   int    `json:"qty"`
	Price string `json:"price"`
	Logo  string `json:"logo"`
}
type Metadata struct {
	MyMetaKey       string `json:"my-meta-key"`
	SettlementGroup string `json:"SettlementGroup"`
}


// -- Create Instapay Payload
type CreateInstapayOrPaymentLinkRequestPayload struct {
  Amount        string `json:"amount"`
  Currency      string `json:"currency"`
  OrderRefID    string `json:"order_ref_id"`
  IsPaymentLink bool   `json:"is_payment_link"`
  Customer      struct {
    Email string `json:"email"`
  } `json:"customer"`
}
