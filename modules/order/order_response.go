package order

import "time"

type CreateOrderResponse struct {
  Data struct {
    ID            string    `json:"id"`
    CustomerID    string    `json:"customer_id"`
    OrderRefID    string    `json:"order_ref_id"`
    Amount        string    `json:"amount"`
    PaymentOption string    `json:"payment_option"`
    Currency      string    `json:"currency"`
    Status        string    `json:"status"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
    ExpiryDate    time.Time `json:"expiry_date"`
    Metadata      struct {
      MyMetaKey       string `json:"my-meta-key"`
      SettlementGroup string `json:"SettlementGroup"`
    } `json:"metadata"`
    Items []struct {
      Name  string `json:"name"`
      Qty   int    `json:"qty"`
      Price string `json:"price"`
      Logo  string `json:"logo"`
    } `json:"items"`
    AccessToken string `json:"access_token"`
  } `json:"data"`
}

// -- Fetch Order Response
type FetchOrderResponse struct {
  Total  int `json:"total"`
  Orders []struct {
    ID            string        `json:"id"`
    Amount        string        `json:"amount"`
    PaymentOption string        `json:"payment_option"`
    Currency      string        `json:"currency"`
    Status        string        `json:"status"`
    OrderRefID    string        `json:"order_ref_id"`
    CreatedAt     int           `json:"created_at"`
    Metadata      []interface{} `json:"metadata"`
  } `json:"orders"`
}

// -- fetch order By Id
type FetchOrderByIdResponse struct {
  ID            string        `json:"id"`
  Amount        string        `json:"amount"`
  PaymentOption string        `json:"payment_option"`
  Currency      string        `json:"currency"`
  Status        string        `json:"status"`
  IsLive        bool          `json:"is_live"`
  OrderRefID    string        `json:"order_ref_id"`
  CreatedAt     int           `json:"created_at"`
  Metadata      []interface{} `json:"metadata"`
}

// -- Create instapay or payment link
type CreateInstapayOrPaymentLinkResponse struct {
  Data struct {
    ID             string        `json:"id"`
    Amount         string        `json:"amount"`
    CustomerID     string        `json:"customer_id"`
    Currency       string        `json:"currency"`
    Status         string        `json:"status"`
    OrderRefID     string        `json:"order_ref_id"`
    CreatedAt      int           `json:"created_at"`
    PaymentLinkURL string        `json:"payment_link_url"`
    Metadata       []interface{} `json:"metadata"`
    AccessToken    string        `json:"access_token"`
  } `json:"data"`
}
