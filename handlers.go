package main

import (
	"net/http"
	"github.com/arschles/serf-proxy/server"
)

func (baseHandler *BaseHandler) keysHandler(resp http.ResponseWriter, req *http.Request) {
	var totalKeys map[string]int
	var totalNumKeys int
	var keys map[string]string
	var bytes []byte
	steps := []*server.HttpStep {
		&server.HttpStep {
			Runner: func() error {
				totKeys, totNumKeys, k, kErr := baseHandler.client.ListKeys()
				if kErr != nil {
					return kErr
				}
				totalKeys = totKeys
				totalNumKeys = totNumKeys
				keys = k
				return nil
			},
			FailCode: http.StatusInternalServerError,
			FailMsg: server.JsonErr(),
		},
		&server.HttpStep {
			Runner: server.EncodeJson(map[string]interface{}{"keys": keys}, &bytes),
			FailCode: http.StatusInternalServerError,
			FailMsg: server.JsonErr(),
		},
	}
	server.NewFailHttpImmediately(resp, steps, http.StatusOK, bytes).Execute()
}

func (baseHandler BaseHandler) statsHandler(resp http.ResponseWriter, req *http.Request) {
	var stats map[string]map[string]string
	var bytes []byte
	steps := []*server.HttpStep {
		&server.HttpStep {
			Runner: func() error {
				s, err := baseHandler.client.Stats()
				if err != nil {
					return err
				}
				stats = s
				return nil
			},
			FailCode: http.StatusInternalServerError,
			FailMsg: server.JsonErr(),
		},
		&server.HttpStep {
			Runner: server.EncodeJson(stats, &bytes),
			FailCode: http.StatusInternalServerError,
			FailMsg: server.JsonErr(),
		},
	}
	server.NewFailHttpImmediately(resp, steps, http.StatusOK, bytes).Execute()
}

func (baseHandler *BaseHandler) useKeyHandler(resp http.ResponseWriter, req *http.Request) {
	var key string
	var keyRing map[string]string
	var bytes []byte
	steps := []*server.HttpStep {
		&server.HttpStep {
			Runner: server.QueryString(req, "key", 0, &key),
			FailCode: http.StatusBadRequest,
			FailMsg: server.JsonErr(),
		},
		&server.HttpStep {
			Runner: func() error {
				k, err := baseHandler.client.UseKey(key)
				if err != nil {
					return err
				}
				keyRing = k
				return nil
			},
			FailCode: http.StatusInternalServerError,
			FailMsg: server.JsonErr(),
		},
		&server.HttpStep {
			Runner: server.EncodeJson(keyRing, &bytes),
			FailCode: http.StatusInternalServerError,
			FailMsg: server.JsonErr(),
		},
	}
	server.NewFailHttpImmediately(resp, steps, http.StatusOK, bytes).Execute()
}
