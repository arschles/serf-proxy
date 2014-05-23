package main

import json_encoder "encoding/json"
import "net/http"
import "fmt"

func writeJson(code int, i interface{}, resp http.ResponseWriter) {
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
