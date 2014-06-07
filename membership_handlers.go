package main

import (
	"encoding/json"
	"fmt"
	SerfClient "github.com/hashicorp/serf/client"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (baseHandler *BaseHandler) deleteMembershipHandler(resp http.ResponseWriter, req *http.Request) {
	err := baseHandler.client.Leave()
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	resp.WriteHeader(http.StatusNoContent)
}

func (baseHandler *BaseHandler) forceDeleteMembershipHandler(resp http.ResponseWriter, req *http.Request) {
	node, err := queryString(req, "node", 0)
	if err != nil {
		writeJsonErr(http.StatusBadRequest, fmt.Errorf("no node in query string"), resp)
		return
	}
	err = baseHandler.client.ForceLeave(node)
	if err != nil {
		writeJsonErr(http.StatusInternalServerError, err, resp)
		return
	}
	resp.WriteHeader(http.StatusNoContent)
}

func (baseHandler *BaseHandler) joinMembershipHandler(resp http.ResponseWriter, req *http.Request) {
	replayStr, err := queryString(req, "replay", 0)
	if err != nil {
		writeJsonErr(http.StatusBadRequest, err, resp)
		return
	}
	replay, err := strconv.ParseBool(replayStr)
	if err != nil {
		writeJsonErr(http.StatusBadRequest, err, resp)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeJsonErr(http.StatusBadRequest, err, resp)
		return
	}

	addrList := []string{}
	err = json.Unmarshal(body, &addrList)
	if err != nil {
		writeJsonErr(http.StatusBadRequest, err, resp)
		return
	}
	i, err := baseHandler.client.Join(addrList, replay)
	if err != nil {
		writeJsonErr(http.StatusInternalServerError, err, resp)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(string(i)))
}

func (baseHandler *BaseHandler) getMembersHandler(resp http.ResponseWriter, req *http.Request) {
	members, err := baseHandler.client.Members()
	if err != nil {
		writeJsonErr(http.StatusInternalServerError, err, resp)
		return
	}
	writeJson(http.StatusOK, map[string][]SerfClient.Member{"members": members}, resp)
}
