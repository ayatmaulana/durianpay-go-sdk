package promo

import "time"

type CreatePromoRequestPayload struct {
	Type         string `json:"type"`
	Label        string `json:"label"`
	Currency     string `json:"currency"`
	PromoDetails struct {
		BinList   []int         `json:"bin_list"`
		BankCodes []interface{} `json:"bank_codes"`
	} `json:"promo_details"`
	DiscountType       string    `json:"discount_type"`
	Discount           string    `json:"discount"`
	MinOrderAmount     string    `json:"min_order_amount"`
	MaxDiscountAmount  string    `json:"max_discount_amount"`
	StartsAt           time.Time `json:"starts_at"`
	EndsAt             time.Time `json:"ends_at"`
	PromoType          string    `json:"promo_type"`
	Description        string    `json:"description"`
	SubType            string    `json:"sub_type"`
	LimitType          string    `json:"limit_type"`
	LimitValue         string    `json:"limit_value"`
	PriceDeductionType string    `json:"price_deduction_type"`
	Code               string    `json:"code"`
}

type UpdatePromoRequestPayload struct {
	Type         string `json:"type"`
	Label        string `json:"label"`
	Currency     string `json:"currency"`
	PromoDetails struct {
		BinList   []int         `json:"bin_list"`
		BankCodes []interface{} `json:"bank_codes"`
	} `json:"promo_details"`
	DiscountType       string    `json:"discount_type"`
	Discount           string    `json:"discount"`
	MinOrderAmount     string    `json:"min_order_amount"`
	MaxDiscountAmount  string    `json:"max_discount_amount"`
	StartsAt           time.Time `json:"starts_at"`
	EndsAt             time.Time `json:"ends_at"`
	PromoType          string    `json:"promo_type"`
	Description        string    `json:"description"`
	SubType            string    `json:"sub_type"`
	LimitType          string    `json:"limit_type"`
	LimitValue         string    `json:"limit_value"`
	PriceDeductionType string    `json:"price_deduction_type"`
	Code               string    `json:"code"`
}
