package main

import (
  "testing"
  . "github.com/franela/goblin"
  "net/http"
)

func TestKeysHandler(t *testing.T) {
  g := Goblin(t)
  g.Describe("keysHandler", func() {
    g.It("should correctly call ListKeys", func() {
      testEnv := NewHandlerTestEnv(DefaultTestEnvURL, EmptyBytes())
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
      testEnv := NewHandlerTestEnv(DefaultTestEnvURL, EmptyBytes())
      g.Assert(testEnv.reqCreateErr).Equal(nil)
      testEnv.baseHandler.statsHandler(testEnv.respRecorder, testEnv.req)
      g.Assert(testEnv.clientTest.StatsCalled).Equal(true)
      g.Assert(testEnv.respRecorder.Code).Equal(http.StatusOK)
    })
  })
}

func TestUseKey(t *testing.T) {
  g := Goblin(t)
  g.Describe("useKeyHandler", func() {
    g.It("should correctly call UseKey", func(){
      testEnv := NewHandlerTestEnv("http://test.com?key=abc", EmptyBytes())
      g.Assert(testEnv.reqCreateErr).Equal(nil)
      testEnv.baseHandler.useKeyHandler(testEnv.respRecorder, testEnv.req)
      g.Assert(testEnv.clientTest.UseKeyCalled).Equal(true)
      g.Assert(testEnv.respRecorder.Code).Equal(http.StatusOK)
    })
  })
}
