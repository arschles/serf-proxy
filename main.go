package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hashicorp/serf/client"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 9999, "the port that the proxy should bind to")
	serfHost := flag.String("serf-host", "localhost", "the host to connect to serf on")
	serfPort := flag.Int("serf-port", 7373, "the port to connect to serf on")
	flag.Parse()

	serfConnStr := fmt.Sprintf("%s:%d", *serfHost, *serfPort)
	log.Printf("connecting to serf on %s", serfConnStr)
	var err error
	serfClient, err := client.NewRPCClient(serfConnStr)
	if err != nil {
		log.Fatalln(err)
	}
	router := mux.NewRouter()

	baseHandler := BaseHandler{client: serfClient}

	//rpcClient.Leave()
	router.HandleFunc("/membership", baseHandler.deleteMembershipHandler).Methods("DELETE")
	//rpcClient.ForceLeave()
	router.HandleFunc("/membership", baseHandler.forceDeleteMembershipHandler).Methods("DELETE").Queries("node", "")
	//rpcClient.Join(addrs, replay)
	router.HandleFunc("/membership", baseHandler.joinMembershipHandler).Methods("POST").Queries("replay", "")
	//rpcClient.Members
	router.HandleFunc("/membership", baseHandler.getMembersHandler).Methods("GET")

	//rpcClient.ListKeys()
	router.HandleFunc("/keys", baseHandler.keysHandler).Methods("GET")

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
	router.HandleFunc("/stats", baseHandler.statsHandler).Methods("GET")

	//rpcClient.Stream
	//TODO

	//rpcClient.UpdateTags
	router.HandleFunc("/tags", baseHandler.updateTagsHandler).Methods("PATCH")

	//rpcClient.UseKey
	router.HandleFunc("/keys", baseHandler.useKeyHandler).Methods("PUT").Queries("key", "")

	//rpcClient.UserEvent
	router.HandleFunc("/event", baseHandler.triggerUserEventHandler).Methods("POST").Queries("coalesce", "")

	log.Printf("serving on port %d", *port)
	addr := fmt.Sprintf(":%d", *port)
	listenAndServeErr := http.ListenAndServe(addr, router)
	if listenAndServeErr != nil {
		log.Fatal(listenAndServeErr)
	}
}
