package refund

import "time"

type CreateRefundResponse struct {
	Data struct {
		ID            string    `json:"id"`
		RefID         string    `json:"ref_id"`
		Amount        string    `json:"amount"`
		RefundType    string    `json:"refund_type"`
		Status        string    `json:"status"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		ApprovedAt    time.Time `json:"approved_at"`
		Source        string    `json:"source"`
		CustomerID    string    `json:"customer_id"`
		CustomerName  string    `json:"customer_name"`
		CustomerEmail string    `json:"customer_email"`
		CustomerPhone string    `json:"customer_phone"`
	} `json:"data"`
}

type FetchRefundResponse struct {
	Data struct {
		Refund []struct {
			ID                 string      `json:"id"`
			MerchantID         string      `json:"merchant_id"`
			Status             string      `json:"status"`
			DisbursementID     interface{} `json:"disbursement_id"`
			TotalAmount        string      `json:"total_amount"`
			CreatedAt          time.Time   `json:"created_at"`
			UpdatedAt          time.Time   `json:"updated_at"`
			ApprovedAt         time.Time   `json:"approved_at"`
			PaymentID          string      `json:"payment_id"`
			RefundRefID        interface{} `json:"refund_ref_id"`
			IsLive             bool        `json:"is_live"`
			Type               string      `json:"type"`
			OrderID            string      `json:"order_id"`
			CustomerID         string      `json:"customer_id"`
			RefundPartial      string      `json:"refund_partial"`
			RefundNotes        string      `json:"refund_notes"`
			PaymentPaidAmount  string      `json:"payment_paid_amount"`
			PaymentDetailsType string      `json:"payment_details_type"`
			PaymentMethodID    string      `json:"payment_method_id"`
			CustomerName       string      `json:"customer_name"`
			CustomerEmail      string      `json:"customer_email"`
			CustomerPhone      string      `json:"customer_phone"`
			Destination        interface{} `json:"destination"`
			EwalletName        interface{} `json:"ewallet_name"`
			AccountName        interface{} `json:"account_name"`
			AccountNumber      interface{} `json:"account_number"`
			CreatedBy          int         `json:"created_by"`
			CreatedByName      string      `json:"created_by_name"`
			UpdatedBy          int         `json:"updated_by"`
			UpdatedByName      string      `json:"updated_by_name"`
			History            interface{} `json:"history"`
			Source             string      `json:"source"`
			AllowRetrigger     bool        `json:"allow_retrigger"`
		} `json:"refund"`
		TotalData int `json:"total_data"`
	} `json:"data"`
}

type FetchrefundByIdResponse struct {
	Data struct {
		ID            string      `json:"id"`
		RefID         interface{} `json:"ref_id"`
		Amount        string      `json:"amount"`
		RefundType    string      `json:"refund_type"`
		Status        string      `json:"status"`
		CreatedAt     time.Time   `json:"created_at"`
		UpdatedAt     time.Time   `json:"updated_at"`
		ApprovedAt    time.Time   `json:"approved_at"`
		Source        string      `json:"source"`
		CustomerID    string      `json:"customer_id"`
		CustomerName  string      `json:"customer_name"`
		CustomerEmail string      `json:"customer_email"`
		CustomerPhone string      `json:"customer_phone"`
	} `json:"data"`
}
