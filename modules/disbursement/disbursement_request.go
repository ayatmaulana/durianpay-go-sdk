package disbursement

type TopupRequestPayload struct {
  BankID int    `json:"bank_id"`
  Amount string `json:"amount"`
}

// -- Submit Disbursement Payload
type Item struct {
    AccountOwnerName string `json:"account_owner_name"`
    BankCode         string `json:"bank_code"`
    Amount           string `json:"amount"`
    AccountNumber    string `json:"account_number"`
    EmailRecipient   string `json:"email_recipient"`
    PhoneNumber      string `json:"phone_number"`
    Notes            string `json:"notes"`
}

type SubmitDisbursementRequestPayload struct {
  Name        string `json:"name"`
  Description string `json:"description"`
  Items       []Item `json:"items"`
}

// -- Validate Bank AccountNumber
type ValidateBankACcountRequestPayload struct {
  AccountNumber string `json:"account_number"`
  BankCode      string `json:"bank_code"`
}
