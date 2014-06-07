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

func (baseHandler *BaseHandler) updateTagsHandler(resp http.ResponseWriter, req *http.Request) {
	//TODO
}

func (baseHandler *BaseHandler) useKeyHandler(resp http.ResponseWriter, req *http.Request) {
	possibleKeys := req.URL.Query()["key"]
	if len(possibleKeys) <= 0 {
		writeJson(http.StatusBadRequest, "invalid key", resp)
		return
	}
	key := possibleKeys[0]
	keyRing, err := baseHandler.client.UseKey(key)
	if err != nil {
		writeJsonErr(http.StatusInternalServerError, err, resp)
		return
	}
	writeJson(http.StatusOK, keyRing, resp)
}
