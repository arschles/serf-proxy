package main

import (
  json_encoder "encoding/json"
  "net/http"
  "fmt"
)

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

func writeJsonErr(code int, err error, resp http.ResponseWriter) {
	jsonStr := fmt.Sprintf(`{"error":"%s"}`, err.Error())
	writeJson(code, jsonStr, resp)
}
