package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func keysHandler(resp http.ResponseWriter, req *http.Request) {
	keys, total, _, err := serfClient.ListKeys()
	if err != nil {
		writeJsonErr(http.StatusInternalServerError, err, resp)
		return
	}
	jsonMap := map[string]interface{}{"keys": keys, "total": total}
	writeJson(http.StatusOK, jsonMap, resp)
}

func statsHandler(resp http.ResponseWriter, req *http.Request) {
	stats, err := serfClient.Stats()
	if err != nil {
		writeJson(http.StatusInternalServerError, err.Error, resp)
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
		writeJson(http.StatusBadRequest, "invalid key", resp)
		return
	}
	keyRing, err := serfClient.UseKey(key)
	if err != nil {
		writeJsonErr(http.StatusInternalServerError, err, resp)
		return
	}
	writeJson(http.StatusOK, keyRing, resp)
}
