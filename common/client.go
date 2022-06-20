package common

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Agent struct {
  ClientConfig *ClientConfig
  HttpClient *http.Client
  Logger *logrus.Logger
}

func (agent *Agent) Call(
  ctx context.Context, 
  method,
  url string, 
  body io.Reader,
) {
  fullUrl := BASE_URL + url

  req, err := http.NewRequest(method, fullUrl, body)

  if err != nil {

  }

  req.WithContext(ctx)
  req.SetBasicAuth(agent.ClientConfig.ApiKey, "")
  req.Header.Add("Content-Type", "application/json")

  res, err := agent.HttpClient.Do(req)

  if err != nil {

  }

  defer res.Body.Close()

  // decode body
  if res.StatusCode >= 200 && res.StatusCode <= 209 {

  } else if res.StatusCode >= 400 && res.StatusCode <= 409 {

  } else if res.StatusCode >= 500 && res.StatusCode <= 509 {

  }


}

func NewAgent(clientConfig *ClientConfig) *Agent {
  return &Agent{
    ClientConfig: clientConfig,
    HttpClient: &http.Client{},
    Logger: logrus.New(),
  }
}
