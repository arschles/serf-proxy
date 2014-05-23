package main

import (
	"net/http"
  "github.com/gorilla/mux"
)

func keysHandler(resp http.ResponseWriter, req *http.Request) {
  keys, total, _, err := serfClient.ListKeys()
  if err != nil {
    http.Error(resp, err.Error(), http.StatusInternalServerError)
    return
  }
  jsonMap := map[string]interface{} {"keys": keys,"total": total}
  writeJson(http.StatusOK, jsonMap, resp)
}

func statsHandler(resp http.ResponseWriter, req *http.Request) {
  stats, err := serfClient.Stats()
  if err != nil {
    writeJson(http.StatusInternalServerError, map[string]string{"error": err.Error()}, resp)
    return
  }
  writeJson(http.StatusOK, stats, resp)
}

func updateTagsHandler(resp http.ResponseWriter, req *http.Request) {
  //TODO
}

func useKeyHandler(resp http.ResponseWriter, req *http.Request) {
  key := mux.Vars(req)["key"]
  if key == "" {
    http.Error(resp, "invalid key", http.StatusBadRequest)
    return
  }
  keyRing, err := serfClient.UseKey(key)
  if err != nil {
    http.Error(resp, err.Error(), http.StatusInternalServerError)
    return
  }
  writeJson(http.StatusOK, keyRing, resp)
}
