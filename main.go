package main

import (
  "github.com/hashicorp/serf/client"
  "flag"
  "fmt"
  "log"
  "github.com/gorilla/mux"
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

  router.HandleFunc("/is_closed", isClosedHandler).Methods("GET")
  router.HandleFunc("/members", membersHandler).Methods("GET")
  router.HandleFunc("/stream", streamHandler).Methods("GET")

  log.Printf("serving on port %d", *port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}
