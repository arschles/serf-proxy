package main

import (
  "github.com/arschles/serf-proxy/test"
  "net/http"
  "net/http/httptest"
  "bytes"
)

const DefaultTestEnvURL = "http://test.com"
func EmptyBytes() []byte {
  return []byte{}
}

type HandlerTestEnv struct {
  clientTest *test.ClientTest
  baseHandler *BaseHandler
  req *http.Request
  reqCreateErr error
  respRecorder *httptest.ResponseRecorder
}

func NewHandlerTestEnv(url string, body []byte) HandlerTestEnv {
  client := test.NewClientTest()
  baseHandler := NewBaseHandler(client)
  reqBody := bytes.NewReader(body)
  req, err := http.NewRequest("GET", url, reqBody)
  respRecorder := httptest.NewRecorder()

  return HandlerTestEnv {
    clientTest: client,
    baseHandler: baseHandler,
    req: req,
    reqCreateErr: err,
    respRecorder: respRecorder,
  }
}
