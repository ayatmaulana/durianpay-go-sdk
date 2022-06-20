package disbursement

import(
  "time"
)

type ResponseBankList struct {
  Data []struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Code        string `json:"code"`
    Instruction struct {
    } `json:"instruction"`
    Type      string    `json:"type"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
  } `json:"data"`
}

// --- Topup
type TopupResponse struct {
  Data struct {
    SenderBank  string    `json:"sender_bank"`
    TotalAmount string    `json:"total_amount"`
    Status      string    `json:"status"`
    ExpiryDate  time.Time `json:"expiry_date"`
    TransferTo  struct {
      BankCode          string `json:"bank_code"`
      BankName          string `json:"bank_name"`
      AtmBersamaCode    string `json:"atm_bersama_code"`
      BankAccountNumber string `json:"bank_account_number"`
      AccountHolderName string `json:"account_holder_name"`
      UniqueCode        int    `json:"unique_code"`
    } `json:"transfer_to"`
  } `json:"data"`
}

// -- Topup History
type TopupHistoryResponse struct {
  Data struct {
    TotalAmount       string    `json:"total_amount"`
    UniqueCode        int       `json:"unique_code"`
    ExpiryDate        time.Time `json:"expiry_date"`
    BankAccountNumber string    `json:"bank_account_number"`
    BankName          string    `json:"bank_name"`
    BankHolderName    string    `json:"bank_holder_name"`
    Status            string    `json:"status"`
  } `json:"data"`
}

// -- Topup Balance
type FetchBalanceResponse struct {
  Data struct {
    Balance int64 `json:"balance"`
  } `json:"data"`
}

// -- Submit 
type SubmitDisbursementResponse struct {
  Data struct {
    ID                 string `json:"id"`
    Name               string `json:"name"`
    TotalAmount        string `json:"total_amount"`
    TotalDisbursements int    `json:"total_disbursements"`
    Description        string `json:"description"`
  } `json:"data"`
}

// -- Approve Disbursement
type ApproveDisbursementResponse struct {
  Data struct {
    ID                 string `json:"id"`
    Name               string `json:"name"`
    Type               string `json:"type"`
    Status             string `json:"status"`
    TotalAmount        string `json:"total_amount"`
    TotalDisbursements int    `json:"total_disbursements"`
    Description        string `json:"description"`
  } `json:"data"`
}

// -- Fetch Disbursementt Items by Id
type FetchDisbursementItemsByIdResponse struct {
  Data struct {
    SubmissionStatus       string `json:"submission_status"`
    DisbursementBatchItems []struct {
      ID                  string `json:"id"`
      DisbursementBatchID string `json:"disbursement_batch_id"`
      AccountOwnerName    string `json:"account_owner_name"`
      RealName            string `json:"real_name"`
      BankCode            string `json:"bank_code"`
      Amount              string `json:"amount"`
      AccountNumber       string `json:"account_number"`
      EmailRecipient      string `json:"email_recipient"`
      PhoneNumber         string `json:"phone_number"`
      InvalidFields       []struct {
        Key     string `json:"key"`
        Message string `json:"message"`
      } `json:"invalid_fields,omitempty"`
      Status    string    `json:"status"`
      Notes     string    `json:"notes"`
      IsDeleted bool      `json:"is_deleted"`
      CreatedAt time.Time `json:"created_at"`
      UpdatedAt time.Time `json:"updated_at"`
    } `json:"disbursement_batch_items"`
    Count int `json:"count"`
  } `json:"data"`
}

// - Fetch Disbursement By IsDeleted
type FetchDisbursementByIdResponse struct {
  Data struct {
    ID                 string    `json:"id"`
    Name               string    `json:"name"`
    Type               string    `json:"type"`
    Status             string    `json:"status"`
    TotalAmount        string    `json:"total_amount"`
    TotalDisbursements int       `json:"total_disbursements"`
    Description        string    `json:"description"`
    Fees               int       `json:"fees"`
    CreatedAt          time.Time `json:"created_at"`
  } `json:"data"`
}

// - Delete Disbursement
type DeleteDisbursementResponse struct {
  Data string `json:"data"`
}

// - Validate Bank AccountNumber
type ValidateBankAccountResponse struct {
  Data struct {
    AccountNumber string `json:"account_number"`
    BankCode      string `json:"bank_code"`
    AccountHolder string `json:"account_holder"`
    Status        string `json:"status"`
  } `json:"data"`
}
