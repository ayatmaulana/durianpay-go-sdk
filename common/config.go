package common

type Mode string

type ClientConfig struct {
  Mode Mode
  ApiKey string
  EnableLogging bool
}
const (
  LIVE Mode = "LIVE"
  SANDBOX Mode = "SANDBOX"

  // --------
  BASE_URL string = "https://api.durianpay.id/v1"


)
