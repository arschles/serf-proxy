package main

import (
	"net/http"
)

func (baseHandler *BaseHandler) keysHandler(resp http.ResponseWriter, req *http.Request) {
	keys, total, _, err := baseHandler.client.ListKeys()
	if err != nil {
		writeJsonErr(http.StatusInternalServerError, err, resp)
		return
	}
	jsonMap := map[string]interface{}{"keys": keys, "total": total}
	writeJson(http.StatusOK, jsonMap, resp)
}

func (baseHandler BaseHandler) statsHandler(resp http.ResponseWriter, req *http.Request) {
	stats, err := baseHandler.client.Stats()
	if err != nil {
		writeJson(http.StatusInternalServerError, err.Error, resp)
		return
	}
	writeJson(http.StatusOK, stats, resp)
}

func (baseHandler *BaseHandler) useKeyHandler(resp http.ResponseWriter, req *http.Request) {
	key, err := queryString(req, "key", 0)
	if err != nil {
		writeJson(http.StatusBadRequest, "invalid key", resp)
		return
	}
	keyRing, err := baseHandler.client.UseKey(key)
	if err != nil {
		writeJsonErr(http.StatusInternalServerError, err, resp)
		return
	}
	writeJson(http.StatusOK, keyRing, resp)
}
