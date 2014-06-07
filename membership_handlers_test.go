package main

import (
  "testing"
  . "github.com/franela/goblin"
  "net/http"
  "encoding/json"
)

func TestDeleteMembershipHandler(t *testing.T) {
  g := Goblin(t)
  g.Describe("testDeleteMembershipHandler", func(){
    g.It("should correctly call Leave", func() {
      testEnv := NewHandlerTestEnv(DefaultTestEnvURL, EmptyBytes())
      g.Assert(testEnv.reqCreateErr).Equal(nil)
      testEnv.baseHandler.deleteMembershipHandler(testEnv.respRecorder, testEnv.req)
      g.Assert(testEnv.clientTest.LeaveCalled).Equal(true)
      g.Assert(testEnv.respRecorder.Code).Equal(http.StatusNoContent)
    })
  })
}

func TestForceDeleteMembershipHandler(t *testing.T) {
  g := Goblin(t)
  g.Describe("testDeleteMembershipHandler", func() {
    g.It("should correctly call ForceLeave", func() {
      testEnv := NewHandlerTestEnv("http://test.com?node=abc", EmptyBytes())
      g.Assert(testEnv.reqCreateErr).Equal(nil)
      testEnv.baseHandler.forceDeleteMembershipHandler(testEnv.respRecorder, testEnv.req)
      g.Assert(testEnv.clientTest.ForceLeaveCalled).Equal(true)
      g.Assert(testEnv.respRecorder.Code).Equal(http.StatusNoContent)
    })
  })
}

func TestJoinMembershipHandler(t *testing.T) {
  g := Goblin(t)
  g.Describe("testDeleteMembershipHandler", func() {
    g.It("should correctly call Join", func() {
      addrList := []string{"hello", "world"}
      addrListBytes, err := json.Marshal(addrList)
      g.Assert(err).Equal(nil)
      testEnv := NewHandlerTestEnv("http://test.com?replay=true", addrListBytes)
      g.Assert(testEnv.reqCreateErr).Equal(nil)

      testEnv.baseHandler.joinMembershipHandler(testEnv.respRecorder, testEnv.req)
      g.Assert(testEnv.clientTest.JoinCalled).Equal(true)
      g.Assert(testEnv.respRecorder.Code).Equal(http.StatusOK)
      g.Assert(len(testEnv.respRecorder.Body.Bytes())).Equal(1)
    })
  })
}

func TestGetMembersHandler(t *testing.T) {
  g := Goblin(t)
  g.Describe("testDeleteMembershipHandler", func() {
    g.It("should correctly call Members", func() {
      testEnv := NewHandlerTestEnv(DefaultTestEnvURL, EmptyBytes())
      g.Assert(testEnv.reqCreateErr).Equal(nil)
      testEnv.baseHandler.getMembersHandler(testEnv.respRecorder, testEnv.req)
      g.Assert(testEnv.clientTest.MembersCalled).Equal(true)
      g.Assert(testEnv.respRecorder.Code).Equal(http.StatusOK)
      //TODO: check resp body
    })
  })
}
