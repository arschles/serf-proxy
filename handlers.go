package main

import (
  "net/http"
  json_encoder "encoding/json"
  "github.com/hashicorp/serf/client"
  "fmt"
)

func json(code int, i interface{}, resp http.ResponseWriter) {
  bytes, err := json_encoder.Marshal(i)
  if err != nil {
    resp.WriteHeader(http.StatusInternalServerError)
    str := fmt.Sprintf(`{"error":""%s"}`, err.Error())
    resp.Write([]byte(str))
  } else {
    resp.WriteHeader(code)
    resp.Write(bytes)
  }
}

func isClosedHandler(resp http.ResponseWriter, req *http.Request) {
  m := map[string]bool{"is_closed":serfClient.IsClosed()}
  json(http.StatusOK, m, resp)
}

func membersHandler(resp http.ResponseWriter, req *http.Request) {
  members, err := serfClient.Members()
  if err != nil {
    json(http.StatusInternalServerError, map[string]string{"error":err.Error()}, resp)
  } else {
    json(http.StatusOK, map[string][]client.Member{"members":members}, resp)
  }
}
