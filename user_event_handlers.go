package main

import "net/http"
import "strconv"
import "encoding/json"
import "io/ioutil"
import "github.com/gorilla/mux"

type userEventRequestPayload struct {
  name string
  payload []byte
}

func triggerUserEventHandler(resp http.ResponseWriter, req *http.Request) {
  coalesce, err := strconv.ParseBool(mux.Vars(req)["coalesce"])
  if err != nil {
    http.Error(resp, "invalid coalesce flag", http.StatusBadRequest)
    return
  }

  body, readAllErr := ioutil.ReadAll(req.Body)
  if readAllErr != nil {
    http.Error(resp, readAllErr.Error(), http.StatusBadRequest)
    return
  }

  payload := userEventRequestPayload{}
  userEventRequestPayloadParseErr := json.Unmarshal(body, &payload)
  if userEventRequestPayloadParseErr != nil {
    http.Error(resp, userEventRequestPayloadParseErr.Error(), http.StatusInternalServerError)
    return
  }
  userEventErr := serfClient.UserEvent(payload.name, payload.payload, coalesce)
  if userEventErr != nil {
    http.Error(resp, userEventErr.Error(), http.StatusInternalServerError)
    return
  }
  resp.WriteHeader(http.StatusNoContent)
}
