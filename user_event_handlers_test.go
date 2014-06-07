package main

import (
  "testing"
  . "github.com/franela/goblin"
  "encoding/json"
  "net/http"
)

func TestTriggerUserEventHandler(t *testing.T) {
  g := Goblin(t)
  g.Describe("triggerUserEventHandler", func() {
    g.It("should correctly call the handler", func() {
      payload := userEventRequestPayload{name:"testN",payload:[]byte("testP")}
      body, err := json.Marshal(payload)
      g.Assert(err).Equal(nil)
      testEnv := NewHandlerTestEnv("http://test.com?coalesce=true", body)
      g.Assert(testEnv.reqCreateErr).Equal(nil)
      testEnv.baseHandler.triggerUserEventHandler(testEnv.respRecorder, testEnv.req)
      g.Assert(testEnv.clientTest.UserEventCalled).Equal(true)
      g.Assert(testEnv.respRecorder.Code).Equal(http.StatusNoContent)
    })
  })
}
