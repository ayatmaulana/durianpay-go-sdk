package common

type ErrorResponse struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}
