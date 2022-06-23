package common

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"net/http"
	"net/http/httputil"

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
  body any,
  response any,
) ( err error) {
  fullUrl := BASE_URL + url
  
  var bodyBuffer bytes.Buffer
  err = json.NewEncoder(&bodyBuffer).Encode(body)

  req, err := http.NewRequest(method, fullUrl, &bodyBuffer)

  if err != nil {
    return
  }

  req.WithContext(ctx)
  req.SetBasicAuth(agent.ClientConfig.ApiKey, "")
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Accept", "application/json")

  res, err := agent.HttpClient.Do(req)

  if err != nil {
    return
  }

  defer res.Body.Close()

  if agent.ClientConfig.EnableLogging {
    agent.debug(req, res)
  }

  // decode body
  if res.StatusCode >= 200 && res.StatusCode <= 209 {
    err = json.NewDecoder(res.Body).Decode(&response)
  } else if res.StatusCode >= 400 && res.StatusCode <= 409 {
    b, err  := io.ReadAll(res.Body)
    if err != nil {
      return err

    }

    agent.Logger.Errorln(string(b))
    agent.Logger.Errorln(err)

  } else if res.StatusCode >= 500 && res.StatusCode <= 509 {
    b, err  := io.ReadAll(res.Body)
    if err != nil {
      return err
    }

    agent.Logger.Errorln(string(b))
    agent.Logger.Errorln(err)
  }

  return
}

func (agent *Agent) debug (req *http.Request, res *http.Response) {
  debugReq, err := httputil.DumpRequest(req, true)
  agent.Logger.Infoln(string(debugReq))

  debug, err := httputil.DumpResponse(res, true)
  if err != nil {}

  agent.Logger.Infoln(string(debug))
}

func NewAgent(clientConfig *ClientConfig) *Agent {
  return &Agent{
    ClientConfig: clientConfig,
    HttpClient: &http.Client{},
    Logger: logrus.New(),
  }
}
