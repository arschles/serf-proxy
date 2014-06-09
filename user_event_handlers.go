package main

import (
	"net/http"
	"github.com/arschles/serf-proxy/server"
)

type userEventRequestPayload struct {
	name    string
	payload []byte
}

func (baseHandler *BaseHandler) triggerUserEventHandler(resp http.ResponseWriter, req *http.Request) {
	var coalesce bool
	var payload userEventRequestPayload
	steps := []*server.HttpStep {
		&server.HttpStep {
			Runner: server.QueryStringParseBool(req, "coalesce", 0, &coalesce),
			FailCode: http.StatusBadRequest,
			FailMsg: server.JsonErr(),
		},
		&server.HttpStep {
			Runner: server.ReadJson(req, &payload),
			FailCode: http.StatusBadRequest,
			FailMsg: server.JsonErr(),
		},
		&server.HttpStep {
			Runner: func() error {
				return baseHandler.client.UserEvent(payload.name, payload.payload, coalesce)
			},
			FailCode: http.StatusInternalServerError,
			FailMsg: server.JsonErr(),
		},
	}
	server.NewFailHttpImmediately(resp, steps, http.StatusNoContent, server.EmptyBody()).Execute()
}
