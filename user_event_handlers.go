package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type userEventRequestPayload struct {
	name    string
	payload []byte
}

func (baseHandler *BaseHandler) triggerUserEventHandler(resp http.ResponseWriter, req *http.Request) {
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
	userEventErr := baseHandler.client.UserEvent(payload.name, payload.payload, coalesce)
	if userEventErr != nil {
		http.Error(resp, userEventErr.Error(), http.StatusInternalServerError)
		return
	}
	resp.WriteHeader(http.StatusNoContent)
}
