package main

import (
	SerfClient "github.com/hashicorp/serf/client"
	"net/http"
	"github.com/arschles/serf-proxy/server"
)

func (baseHandler *BaseHandler) deleteMembershipHandler(resp http.ResponseWriter, req *http.Request) {
	steps := []*server.HttpStep {
		&server.HttpStep {
			Runner: func() error {
				return baseHandler.client.Leave()
			},
			FailCode: http.StatusInternalServerError,
			FailMsg: server.JsonErr(),
		},
		&server.HttpStep {
			Runner: func() error { return baseHandler.client.Leave() },
			FailCode: http.StatusInternalServerError,
			FailMsg: server.JsonErr(),
		},
	}
	server.NewFailHttpImmediately(resp, steps, http.StatusNoContent, server.EmptyBody()).Execute()
}

func (baseHandler *BaseHandler) forceDeleteMembershipHandler(resp http.ResponseWriter, req *http.Request) {
	var node string

	steps := []*server.HttpStep {
		&server.HttpStep {
			Runner: server.QueryString(req, "node", 0, &node),
			FailCode: http.StatusBadRequest,
			FailMsg: server.JsonErr(),
		},
		&server.HttpStep {
			Runner: func() error {
				return baseHandler.client.ForceLeave(node)
			},
			FailCode: http.StatusInternalServerError,
			FailMsg: server.JsonErr(),
		},
	}
	server.NewFailHttpImmediately(resp, steps, http.StatusNoContent, server.EmptyBody()).Execute()
}

func (baseHandler *BaseHandler) joinMembershipHandler(resp http.ResponseWriter, req *http.Request) {
	var replay bool
	var addrList []string
	var joinNum int
	steps := []*server.HttpStep {
		&server.HttpStep {
			Runner: server.QueryStringParseBool(req, "replay", 0, &replay),
			FailCode: http.StatusBadRequest,
			FailMsg: server.JsonErr(),
		},
		&server.HttpStep {
			Runner: server.ReadJson(req, &addrList),
			FailCode: http.StatusBadRequest,
			FailMsg: server.JsonErr(),
		},
		&server.HttpStep {
			Runner: func() error {
				i, err := baseHandler.client.Join(addrList, replay)
				if err != nil {
					return err
				}
				joinNum = i
				return nil
			},
			FailCode: http.StatusInternalServerError,
			FailMsg: server.JsonErr(),
		},
	}
	server.NewFailHttpImmediately(resp, steps, http.StatusOK, []byte(string(joinNum))).Execute()
}

func (baseHandler *BaseHandler) getMembersHandler(resp http.ResponseWriter, req *http.Request) {
	var members []SerfClient.Member
	var bytes []byte
	steps := []*server.HttpStep {
		&server.HttpStep {
			Runner: func() error {
				m, err := baseHandler.client.Members()
				if err != nil {
					return err
				}
				members = m
				return nil
			},
			FailCode: http.StatusInternalServerError,
			FailMsg: server.JsonErr(),
		},
		&server.HttpStep {
			Runner: server.EncodeJson(map[string][]SerfClient.Member{"members":members}, &bytes),
			FailCode: http.StatusInternalServerError,
			FailMsg: server.JsonErr(),
		},
	}

	server.NewFailHttpImmediately(resp, steps, http.StatusOK, bytes).Execute()
}
