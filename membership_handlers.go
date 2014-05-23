package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "strconv"
  "encoding/json"
  "io/ioutil"
  SerfClient "github.com/hashicorp/serf/client"
)

func deleteMembershipHandler(resp http.ResponseWriter, req *http.Request) {
  err := serfClient.Leave()
  if err != nil {
    http.Error(resp, err.Error(), http.StatusInternalServerError)
    return
  }
  resp.WriteHeader(http.StatusNoContent)
}

func forceDeleteMembershipHandler(resp http.ResponseWriter, req *http.Request) {
  node := mux.Vars(req)["node"]
  if node == "" {
    http.Error(resp, "no node in query string", http.StatusBadRequest)
    return
  }
  err := serfClient.ForceLeave(node)
  if err != nil {
    http.Error(resp, err.Error(), http.StatusInternalServerError)
    return
  }
  resp.WriteHeader(http.StatusNoContent)
}

func joinMembershipHandler(resp http.ResponseWriter, req *http.Request) {
  replay, replayParseErr := strconv.ParseBool(mux.Vars(req)["replay"])
  if replayParseErr != nil {
    http.Error(resp, replayParseErr.Error(), http.StatusBadRequest)
    return
  }

  body, readAllErr := ioutil.ReadAll(req.Body)
  if readAllErr != nil {
    http.Error(resp, readAllErr.Error(), http.StatusBadRequest)
    return
  }

  addrList := []string{}
  addrListParseErr := json.Unmarshal(body, &addrList)
  if addrListParseErr != nil {
    http.Error(resp, addrListParseErr.Error(), http.StatusBadRequest)
    return
  }
  i, joinErr := serfClient.Join(addrList, replay)
  if joinErr != nil {
    http.Error(resp, joinErr.Error(), http.StatusInternalServerError)
    return
  }

  resp.WriteHeader(http.StatusOK)
  resp.Write([]byte(string(i)))
}

func getMembersHandler(resp http.ResponseWriter, req *http.Request) {
  members, err := serfClient.Members()
  if err != nil {
    writeJson(http.StatusInternalServerError, map[string]string{"error": err.Error()}, resp)
  } else {
    writeJson(http.StatusOK, map[string][]SerfClient.Member{"members": members}, resp)
  }
}
