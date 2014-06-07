package main

import (
  "testing"
  "github.com/arschles/serf-proxy/test"
  . "github.com/franela/goblin"
  "net/http"
  "net/http/httptest"
)

const DefaultTestEnvURL = "http://test.com"

type HandlerTestEnv struct {
  clientTest *test.ClientTest
  baseHandler *BaseHandler
  req *http.Request
  reqCreateErr error
  respRecorder *httptest.ResponseRecorder
}

func setupHandlerTestEnv(url string) HandlerTestEnv {
  client := test.NewClientTest()
  baseHandler := NewBaseHandler(client)
  req, err := http.NewRequest("GET", url, nil)
  respRecorder := httptest.NewRecorder()

  return HandlerTestEnv {
    clientTest: client,
    baseHandler: baseHandler,
    req: req,
    reqCreateErr: err,
    respRecorder: respRecorder,
  }
}

func TestKeysHandler(t *testing.T) {
  g := Goblin(t)
  g.Describe("keysHandler", func() {
    g.It("should correctly call ListKeys", func() {
      testEnv := setupHandlerTestEnv(DefaultTestEnvURL)
      g.Assert(testEnv.reqCreateErr).Equal(nil)
      testEnv.baseHandler.keysHandler(testEnv.respRecorder, testEnv.req)
      g.Assert(testEnv.clientTest.ListKeysCalled).Equal(true)
      g.Assert(testEnv.respRecorder.Code).Equal(http.StatusOK)
    })
  })
}

func TestStatsHandler(t *testing.T) {
  g := Goblin(t)
  g.Describe("statsHandler", func() {
    g.It("should correctly call Stats", func() {
      testEnv := setupHandlerTestEnv(DefaultTestEnvURL)
      g.Assert(testEnv.reqCreateErr).Equal(nil)
      testEnv.baseHandler.statsHandler(testEnv.respRecorder, testEnv.req)
      g.Assert(testEnv.clientTest.StatsCalled).Equal(true)
      g.Assert(testEnv.respRecorder.Code).Equal(http.StatusOK)
    })
  })
}

func TestUpdateTagsHandler(t *testing.T) {
  g := Goblin(t)
  g.Describe("updateTagsHandler", func() {
    g.It("should correctly call UpdateTags", func() {
      //TODO: implement
    })
  })
}

func TestUseKey(t *testing.T) {
  g := Goblin(t)
  g.Describe("useKeyHandler", func() {
    g.It("should correctly call UseKey", func(){
      testEnv := setupHandlerTestEnv("http://test.com?key=abc")
      g.Assert(testEnv.reqCreateErr).Equal(nil)
      testEnv.baseHandler.useKeyHandler(testEnv.respRecorder, testEnv.req)
      g.Assert(testEnv.clientTest.UseKeyCalled).Equal(true)
      g.Assert(testEnv.respRecorder.Code).Equal(http.StatusOK)
    })
  })
}
