package main

import (
	json_encoder "encoding/json"
	"fmt"
	"net/http"
)

const ErrorKey = "error"

type BaseHandler struct {
	client Client
}

func NewBaseHandler(cl Client) *BaseHandler {
	return &BaseHandler{client:cl}
}

func writeJson(code int, i interface{}, resp http.ResponseWriter) {
	bytes, err := json_encoder.Marshal(i)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		str := fmt.Sprintf(`{"error":"»%s"}`, err.Error())
		resp.Write([]byte(str))
	} else {
		resp.WriteHeader(code)
		resp.Write(bytes)
	}
}

func writeJsonErr(code int, err error, resp http.ResponseWriter) {
	errMap := map[string]string{ErrorKey:err.Error()}
	writeJson(code, errMap, resp)
}

//returns the value at idx for key in req's query string.
//returns "" and an appropriate error if key doesn't exist, or there
//aren't enough values under that key. idx is 0-based
func queryString(req *http.Request, key string, idx int) (string, error) {
	queryString := req.URL.Query()
	vals := queryString[key]
	if vals == nil {
		return "", fmt.Errorf("no key %s in the query string", key)
	} else if idx >= len(vals) {
		return "", fmt.Errorf("fewer than %d values for key %s in the query string", idx+1, key)
	}
	return vals[idx], nil
}
