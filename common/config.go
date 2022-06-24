package common

type Mode string

type ClientConfig struct {
  Mode Mode
  ApiKey string
  EnableLogging bool
}

type QueryParams struct {
  from int `url:"from,omitempty"`
  to int `url:"to,omitempty"`
  skip int `url:"skip,omitempty"`
  limit int `url:"limit,omitempty"`
}


const (
  LIVE Mode = "LIVE"
  SANDBOX Mode = "SANDBOX"

  // --------
  BASE_URL string = "https://api.durianpay.id/v1"


)
