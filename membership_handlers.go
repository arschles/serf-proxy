package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	SerfClient "github.com/hashicorp/serf/client"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (baseHandler BaseHandler) deleteMembershipHandler(resp http.ResponseWriter, req *http.Request) {
	err := serfClient.Leave()
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	resp.WriteHeader(http.StatusNoContent)
}

func (baseHandler BaseHandler) forceDeleteMembershipHandler(resp http.ResponseWriter, req *http.Request) {
	node := mux.Vars(req)["node"]
	if node == "" {
		writeJsonErr(http.StatusBadRequest, fmt.Errorf("no node in query string"), resp)
		return
	}
	err := serfClient.ForceLeave(node)
	if err != nil {
		writeJsonErr(http.StatusInternalServerError, err, resp)
		return
	}
	resp.WriteHeader(http.StatusNoContent)
}

func (baseHandler BaseHandler) joinMembershipHandler(resp http.ResponseWriter, req *http.Request) {
	replay, replayParseErr := strconv.ParseBool(mux.Vars(req)["replay"])
	if replayParseErr != nil {
		writeJsonErr(http.StatusBadRequest, replayParseErr, resp)
		return
	}

	body, readAllErr := ioutil.ReadAll(req.Body)
	if readAllErr != nil {
		writeJsonErr(http.StatusBadRequest, readAllErr, resp)
		return
	}

	addrList := []string{}
	addrListParseErr := json.Unmarshal(body, &addrList)
	if addrListParseErr != nil {
		writeJsonErr(http.StatusBadRequest, addrListParseErr, resp)
		return
	}
	i, joinErr := serfClient.Join(addrList, replay)
	if joinErr != nil {
		writeJsonErr(http.StatusInternalServerError, joinErr, resp)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(string(i)))
}

func (baseHandler BaseHandler) getMembersHandler(resp http.ResponseWriter, req *http.Request) {
	members, err := serfClient.Members()
	if err != nil {
		writeJsonErr(http.StatusInternalServerError, err, resp)
		return
	}
	writeJson(http.StatusOK, map[string][]SerfClient.Member{"members": members}, resp)
}
