package payment

type EWALLET string
type BANK string
type RETAIL_STORE string
type ONLINE_BANKING string
type BNPL string // buy now pay later

const (
  DANA EWALLET = "DANA"
  SHOPEEPAY = "SHOPEEPAY"
  LINKAJA = "LINKAJA"
  OVO = "OVO"
  GOPAY = "GOPAY"
)

const (
  BCA BANK = "BCA"
  BNI = "BNI"
  PERMATA = "PERMATA"
  BRI = "BRI"
  MANDIRI = "MANDIRI"
  CIMB = "CIMB"
  BTPN = "BTPN"
  SAHABAT_SAMPOERNA = "SAHABAT_SAMPOERNA"
  SINARMAS = "SINARMAS"
  MAYBANK = "MAYBANK"
  SYARIAH = "SYARIAH"
  DANAMON = "DANAMON"
  OTHER = "OTHER"
)

const (
  ALFAMART RETAIL_STORE = "ALFAMART"
  INDOMARET = "INDOMARET"
)

const (
  JENIUSPAY ONLINE_BANKING = "JENIUSPAY"
)

const (
  INDODANA BNPL = "INDODANA"
  AKULAKU = "AKULAKU"
)

type CustomerInfo struct {
  ID        string `json:"id"`
  Email     string `json:"email"`
  GivenName string `json:"given_name"`
}

type ChargePaymentVARequestPayload struct {
	OrderID  string `json:"order_id"`
	BankCode BANK `json:"bank_code"`
	Name     string `json:"name"`
	Amount   string `json:"amount"`
}

type ChargePaymentEWalletRequestPayload struct {
	OrderID    string `json:"order_id"`
	Amount     string `json:"amount"`
	Mobile     string `json:"mobile"`
	WalletType EWALLET `json:"wallet_type"`
}

type ChargePaymentRetailStorePayload struct {
	OrderID  string `json:"order_id"`
	BankCode RETAIL_STORE `json:"bank_code"`
	Name     string `json:"name"`
	Amount   string `json:"amount"`
}

type ChargePaymentCardRequestPayload struct {
	OrderID      string `json:"order_id"`
	Amount       string `json:"amount"`
        CustomerInfo          CustomerInfo `json:"customer_info"`
}

type ChargePaymentOnlineBankingRequestPayload struct {
	OrderID      string `json:"order_id"`
	Type         ONLINE_BANKING `json:"type"`
	Name         string `json:"name"`
	Amount       string `json:"amount"`
        CustomerInfo          CustomerInfo `json:"customer_info"`
	Mobile string `json:"mobile"`
}

type ChargePaymentQRISRequestPayload struct {
	Amount  string `json:"amount"`
	OrderID string `json:"order_id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

type ChargePaymentBNPLRequestPayload struct {
	OrderID               string `json:"order_id"`
	Amount                string `json:"amount"`
	PaymentMethodUniqueID BNPL `json:"payment_method_unique_id"`
        CustomerInfo          CustomerInfo `json:"customer_info"`
}

