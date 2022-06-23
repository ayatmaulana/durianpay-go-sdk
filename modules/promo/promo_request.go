package promo

import "time"

type TYPE string
type SUBTYPE string
type DISCOUNT_TYPE string
type LIMIT_TYPE string
type PRICE_DEDUCTION_TYPE string

const (
  CARD_PROMO TYPE = "card_promos"
  EWALLET_PROMO TYPE = "ewallet_promos"
  VA_PROMO TYPE = "va_promos"

  DIRECT_DISCOUNT SUBTYPE = "direct_discount"
  CASHBACK SUBTYPE = "cashback"


  PERCENTAGE DISCOUNT_TYPE = "percentage"
  FIAT DISCOUNT_TYPE = "fiat"

  QUOTA LIMIT_TYPE = "quota"
  BUDGET LIMIT_TYPE = "budget"

  TOTAL_PRICE PRICE_DEDUCTION_TYPE = "total_price"
  PRODUCT_PRICE PRICE_DEDUCTION_TYPE = "product_price"
  SHIPPING_PRICE PRICE_DEDUCTION_TYPE = "shipping_price"
)

type PromoDetails struct {
  BinList   []int         `json:"bin_list"`
  BankCodes []interface{} `json:"bank_codes"`
}

type CreatePromoRequestPayload struct {
	Type         TYPE `json:"type"`
	Label        string `json:"label"`
	Currency     string `json:"currency"`
	PromoDetails PromoDetails `json:"promo_details"`
	DiscountType       DISCOUNT_TYPE    `json:"discount_type"`
	Discount           string    `json:"discount"`
	MinOrderAmount     string    `json:"min_order_amount"`
	MaxDiscountAmount  string    `json:"max_discount_amount"`
	StartsAt           time.Time `json:"starts_at"`
	EndsAt             time.Time `json:"ends_at"`
	PromoType          string    `json:"promo_type"`
	Description        string    `json:"description"`
	SubType            SUBTYPE    `json:"sub_type"`
	LimitType          LIMIT_TYPE    `json:"limit_type"`
	LimitValue         string    `json:"limit_value"`
	PriceDeductionType PRICE_DEDUCTION_TYPE    `json:"price_deduction_type"`
	Code               string    `json:"code"`
}

type UpdatePromoRequestPayload struct {
	Type         string `json:"type"`
	Label        string `json:"label"`
	Currency     string `json:"currency"`
        PromoDetails PromoDetails `json:"promo_details"`
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
