package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hashicorp/serf/client"
	"log"
	"net/http"
)

var serfClient *client.RPCClient

func main() {
	port := flag.Int("port", 9999, "the port that the proxy should bind to")
	serfHost := flag.String("serf-host", "localhost", "the host to connect to serf on")
	serfPort := flag.Int("serf-port", 7373, "the port to connect to serf on")
	flag.Parse()

	serfConnStr := fmt.Sprintf("%s:%d", *serfHost, *serfPort)
	log.Printf("connecting to serf on %s", serfConnStr)
	var err error
	serfClient, err = client.NewRPCClient(serfConnStr)
	if err != nil {
		log.Fatalln(err)
	}
	router := mux.NewRouter()

  //rpcClient.Leave()
  router.HandleFunc("/membership", deleteMembershipHandler).Methods("DELETE")
  //rpcClient.ForceLeave()
  router.HandleFunc("/membership", forceDeleteMembershipHandler).Methods("DELETE").Queries("node", "")
  //rpcClient.Join(addrs, replay)
  router.HandleFunc("/membership", joinMembershipHandler).Methods("POST").Queries("replay", "")
  //rpcClient.Members
  router.HandleFunc("/membership", getMembersHandler).Methods("GET")

  //rpcClient.ListKeys()
  router.HandleFunc("/keys", keysHandler).Methods("GET")

  //rpcClient.MembersFiltered
  //TODO

  //rpcClient.Monitor
  //TODO

  //rpcClient.Query
  //TODO

  //rpcClient.RemoveKey
  //TODO

  //rpcClient.Respond
  //TODO

  //rpcClient.Stats
	router.HandleFunc("/stats", statsHandler).Methods("GET")

  //rpcClient.Stream
  //TODO

  //rpcClient.UpdateTags
  router.HandleFunc("/tags", updateTagsHandler).Methods("PATCH")

  //rpcClient.UseKey
  router.HandleFunc("/keys", useKeyHandler).Methods("PUT").Queries("key", "")

  //rpcClient.UserEvent
  router.HandleFunc("/event", triggerUserEventHandler).Methods("POST").Queries("coalesce", "")

	log.Printf("serving on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}
