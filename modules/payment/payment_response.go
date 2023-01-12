package payment

import "time"

type FetchPaymentResponse struct {
	Payments []struct {
		ID           string        `json:"id"`
		Amount       string        `json:"amount"`
		Currency     string        `json:"currency"`
		Status       string        `json:"status"`
		OrderID      string        `json:"order_id"`
		PaymentRefID string        `json:"payment_ref_id"`
		CreatedAt    int           `json:"created_at"`
		Metadata     []interface{} `json:"metadata"`
	} `json:"payments"`
	Total int `json:"total"`
}

type FetchPaymentByIdResponse struct {
	ID           string        `json:"id"`
	Amount       string        `json:"amount"`
	Currency     string        `json:"currency"`
	Status       string        `json:"status"`
	OrderID      string        `json:"order_id"`
	PaymentRefID string        `json:"payment_ref_id"`
	CreatedAt    int           `json:"created_at"`
	Metadata     []interface{} `json:"metadata"`
}

type ChargePaymentVAResponse struct {
	Data struct {
		Type     string `json:"type"`
		Response struct {
			PaymentID          string    `json:"payment_id"`
			OrderID            string    `json:"order_id"`
			AccountNumber      string    `json:"account_number"`
			ExpirationTime     time.Time `json:"expiration_time"`
			PaidAmount         string    `json:"paid_amount"`
			PaymentInstruction struct {
			} `json:"payment_instruction"`
			Metadata struct {
			} `json:"metadata"`
		} `json:"response"`
	} `json:"data"`
}

type ChargePaymentEWalletResponse struct {
  Data struct {
    Type     string `json:"type"`
    Response struct {
      PaymentID      string    `json:"payment_id"`
      OrderID        string    `json:"order_id"`
      Mobile         string    `json:"mobile"`
      Status         string    `json:"status"`
      ExpirationTime time.Time `json:"expiration_time"`
      CheckoutURL    string    `json:"checkout_url"`
      PaidAmount     string    `json:"paid_amount"`
      Metadata       struct {
      } `json:"metadata"`
    } `json:"response"`
  }
}

type ChargePaymentRetailStoreResponse struct {
  Data struct {
    Type     string `json:"type"`
    Response struct {
      PaymentID      string    `json:"payment_id"`
      OrderID        string    `json:"order_id"`
      AccountNumber  string    `json:"account_number"`
      ExpirationTime time.Time `json:"expiration_time"`
      PaidAmount     string    `json:"paid_amount"`
      Metadata       struct {
      } `json:"metadata"`
    } `json:"response"`
  }
}

type ChargePaymentCardResponse struct {
  Data struct {
    Type     string `json:"type"`
    Response struct {
      PaymentID    string `json:"payment_id"`
      OrderID      string `json:"order_id"`
      PaymentRefID string `json:"payment_ref_id"`
      TokenID      string `json:"token_id"`
      Status       string `json:"status"`
      PaidAmount   string `json:"paid_amount"`
      Metadata     struct {
      } `json:"metadata"`
      CheckoutURL string `json:"checkout_url"`
    } `json:"resp}onse"`
  }
}

type ChargePaymentOnlineBankingResponse struct {
  Data struct {
    Type     string `json:"type"`
    Response struct {
      ExpirationTime time.Time `json:"expiration_time"`
      Metadata       struct {
      } `json:"metadata"`
      Mobile     string `json:"mobile"`
      OrderID    string `json:"order_id"`
      PaidAmount string `json:"paid_amount"`
      PaymentID  string `json:"payment_id"`
      Status     string `json:"status"`
      UniqueID   string `json:"unique_id"`
      WebURL     string `json:"web_url"`
    } `json:"response"`
  }
}

type ChargePaymentQRISResponse struct {
  Data struct {
    Type     string `json:"type"`
    Response struct {
      PaymentID      string    `json:"payment_id"`
      OrderID        string    `json:"order_id"`
      Status         string    `json:"status"`
      ExpirationTime time.Time `json:"expiration_time"`
      CreationTime   time.Time `json:"creation_time"`
      QrString       string    `json:"qr_string"`
      UniqueID       string    `json:"unique_id"`
      Metadata       struct {
        MerchantName string `json:"merchant_name"`
        MerchantID   string `json:"merchant_id"`
      } `json:"metadata"`
      Amount string `json:"amount"`
      QrCode string `json:"qr_code"`
    } `json:"response"`
  }
}

type ChargePaymentBNPLResponse struct {
  Data struct {
    Type     string `json:"type"`
    Response struct {
      PaymentID   string `json:"payment_id"`
      OrderID     string `json:"order_id"`
      RedirectURL string `json:"redirect_url"`
      PaidAmount  string `json:"paid_amount"`
      Metadata    struct {
      } `json:"metadata"`
    } `json:"response"`
  }
}

type CheckPaymentStatusResponse struct {
	Data struct {
		Status      string `json:"status"`
		IsCompleted bool   `json:"is_completed"`
		Signature   string `json:"signature"`
	} `json:"data"`
}

type VerifyPaymentResponse struct {
	Data bool `json:"data"`
}

type CancelPaymentResponse struct {
	Data struct {
		ID        string    `json:"id"`
		Status    string    `json:"status"`
		OrderID   string    `json:"order_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"data"`
}

type MDRFeeCalculationResponse struct {
	Data map[string]interface{}
}
