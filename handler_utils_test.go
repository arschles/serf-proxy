package main

import (
  "testing"
  . "github.com/franela/goblin"
  "net/http"
  "net/http/httptest"
  "encoding/json"
  "fmt"
)

func TestWriteJson(t *testing.T) {
  g := Goblin(t)
  g.Describe("writeJson", func() {
    g.It("should properly write the code and encoded json", func() {
      code := http.StatusOK
      toWrite := map[string]string{"hello":"world"}
      respWriter := httptest.NewRecorder()
      writeJson(code, toWrite, respWriter)
      g.Assert(respWriter.Code).Equal(code)
      expectedBytes, err := json.Marshal(toWrite)
      g.Assert(err).Equal(nil)
      g.Assert(string(respWriter.Body.Bytes())).Equal(string(expectedBytes))
    })
  })
}

func TestWriteJsonErr(t *testing.T) {
  g := Goblin(t)
  g.Describe("writeJsonErr", func() {
    g.It("should properly write the code and encoded json", func() {
      code := http.StatusInternalServerError
      err := fmt.Errorf("test error")
      respWriter := httptest.NewRecorder()
      writeJsonErr(code, err, respWriter)
      g.Assert(respWriter.Code).Equal(code)
      expectedErrMap := map[string]string{"error":err.Error()}
      expectedBytes, marshalErr := json.Marshal(expectedErrMap)
      g.Assert(marshalErr).Equal(nil)
      g.Assert(string(respWriter.Body.Bytes())).Equal(string(expectedBytes))
    })
  })
}

func TestQueryString(t *testing.T) {
  g := Goblin(t)
  g.Describe("queryString", func() {
    g.It("should correctly get the right query string value when the key exists", func() {
      req, err := http.NewRequest("GET", "http://test.com?a=b&a=c", nil)
      g.Assert(err).Equal(nil)
      first, err := queryString(req, "a", 0)
      g.Assert(err).Equal(nil)
      g.Assert(first).Equal("b")
      second, err := queryString(req, "a", 1)
      g.Assert(err).Equal(nil)
      g.Assert(second).Equal("c")
    })
    g.It("should correctly fail if there's no key in the query string", func() {
      req, err := http.NewRequest("GET", "http://test.com?a=b", nil)
      g.Assert(err).Equal(nil)
      res, err := queryString(req, "b", 0)
      g.Assert(err != nil).IsTrue()
      g.Assert(res).Equal("")
    })
    g.It("should correctly fail if there are not enough values for the key in the query string", func() {
      req, err := http.NewRequest("GET", "http://test.com?a=b", nil)
      g.Assert(err).Equal(nil)
      res, err := queryString(req, "a", 1)
      g.Assert(err != nil).IsTrue()
      g.Assert(res).Equal("")
    })
  })
}
