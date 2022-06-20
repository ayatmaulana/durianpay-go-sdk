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
