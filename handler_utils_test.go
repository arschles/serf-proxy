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
