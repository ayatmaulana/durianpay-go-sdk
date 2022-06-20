package promo

import "time"


type CreatePromoResponse struct {
	Data struct {
		Currency          string    `json:"currency"`
		Label             string    `json:"label"`
		Description       string    `json:"description"`
		MinOrderAmount    string    `json:"min_order_amount"`
		MaxDiscountAmount string    `json:"max_discount_amount"`
		StartsAt          time.Time `json:"starts_at"`
		EndsAt            time.Time `json:"ends_at"`
		Discount          string    `json:"discount"`
		DiscountType      string    `json:"discount_type"`
		Type              string    `json:"type"`
		PromoDetails      struct {
			PromoID   string        `json:"promo_id"`
			BinList   []int         `json:"bin_list"`
			BankCodes []interface{} `json:"bank_codes"`
		} `json:"promo_details"`
		SubType            string    `json:"sub_type"`
		LimitType          string    `json:"limit_type"`
		LimitValue         string    `json:"limit_value"`
		PriceDeductionType string    `json:"price_deduction_type"`
		Status             string    `json:"status"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		IsLive             bool      `json:"is_live"`
		PromoUsage         string    `json:"promo_usage"`
		ID                 string    `json:"id"`
	} `json:"data"`
}

type FetchPromoResponse struct {
	Data []struct {
		ID                string `json:"id"`
		Currency          string `json:"currency"`
		Label             string `json:"label"`
		Description       string `json:"description"`
		Status            string `json:"status"`
		StartsAt          string `json:"starts_at"`
		EndsAt            string `json:"ends_at"`
		MinOrderAmount    string `json:"min_order_amount"`
		MaxDiscountAmount string `json:"max_discount_amount"`
		Type              string `json:"type"`
		PromoDetails      struct {
			BinList   []int    `json:"bin_list"`
			BankCodes []string `json:"bank_codes"`
		} `json:"promo_details"`
		Discount           int    `json:"discount"`
		DiscountType       string `json:"discount_type"`
		LimitType          string `json:"limit_type"`
		Limit              int    `json:"limit"`
		SubType            string `json:"sub_type"`
		PriceDeductionType string `json:"price_deduction_type"`
	} `json:"data"`
}

type FetchPromoByIdResponse struct {
	ID                string `json:"id"`
	Currency          string `json:"currency"`
	Label             string `json:"label"`
	Description       string `json:"description"`
	Status            string `json:"status"`
	StartsAt          string `json:"starts_at"`
	EndsAt            string `json:"ends_at"`
	MinOrderAmount    string `json:"min_order_amount"`
	MaxDiscountAmount string `json:"max_discount_amount"`
	Type              string `json:"type"`
	PromoDetails      struct {
		BinList   []int    `json:"bin_list"`
		BankCodes []string `json:"bank_codes"`
	} `json:"promo_details"`
	Discount           string `json:"discount"`
	DiscountType       string `json:"discount_type"`
	LimitType          string `json:"limit_type"`
	Limit              int    `json:"limit"`
	SubType            string `json:"sub_type"`
	PriceDeductionType string `json:"price_deduction_type"`
}

type UpdatePromoResponse struct {
	ID                string `json:"id"`
	Currency          string `json:"currency"`
	Label             string `json:"label"`
	Description       string `json:"description"`
	Code              string `json:"code"`
	Status            string `json:"status"`
	StartsAt          string `json:"starts_at"`
	EndsAt            string `json:"ends_at"`
	MinOrderAmount    string `json:"min_order_amount"`
	MaxDiscountAmount string `json:"max_discount_amount"`
	Type              string `json:"type"`
	PromoDetails      struct {
		BinList   []int    `json:"bin_list"`
		BankCodes []string `json:"bank_codes"`
	} `json:"promo_details"`
	Discount           int    `json:"discount"`
	DiscountType       string `json:"discount_type"`
	LimitType          string `json:"limit_type"`
	Limit              int    `json:"limit"`
	SubType            string `json:"sub_type"`
	PriceDeductionType string `json:"price_deduction_type"`
}

type DeletePromoResponse struct {
	Message string `json:"message"`
}
