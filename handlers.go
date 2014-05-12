package main

import (
  "net/http"
  json_encoder "encoding/json"
  "github.com/hashicorp/serf/client"
  "fmt"
  "log"
)

func json(code int, i interface{}, resp http.ResponseWriter) {
  bytes, err := json_encoder.Marshal(i)
  if err != nil {
    resp.WriteHeader(http.StatusInternalServerError)
    str := fmt.Sprintf(`{"error":""%s"}`, err.Error())
    resp.Write([]byte(str))
  } else {
    resp.WriteHeader(code)
    resp.Write(bytes)
  }
}

func isClosedHandler(resp http.ResponseWriter, req *http.Request) {
  m := map[string]bool{"is_closed":serfClient.IsClosed()}
  json(http.StatusOK, m, resp)
}

func membersHandler(resp http.ResponseWriter, req *http.Request) {
  members, err := serfClient.Members()
  if err != nil {
    json(http.StatusInternalServerError, map[string]string{"error":err.Error()}, resp)
  } else {
    json(http.StatusOK, map[string][]client.Member{"members":members}, resp)
  }
}

func streamHandler(resp http.ResponseWriter, req *http.Request) {
  streamCh := make(chan map[string]interface{})
  //TODO: add filter to request
  handle, err := serfClient.Stream("*", streamCh)
  if err != nil {
    json(http.StatusInternalServerError, map[string]string{"error":err.Error()}, resp)
  } else {
    log.Printf("beginning to chunk results from stream %d", handle)
    resp.Header().Set("Transfer-Encoding", "chunked")
    resp.Header().Set("Content-Type", "application/octet-stream")
    resp.Header().Set("Content-Length", "0")
    resp.Write([]byte("hello world"))
    /*
    for streamed := range(streamCh) {
      log.Printf("chunking %+v", streamed)
      bytes, err := json_encoder.Marshal(streamed)
      //TODO: send error to chunked resp
      if err == nil {
        resp.Write(bytes)
      }
      log.Printf("chunking succeeded")
    }
    log.Printf("chunking done for stream %d", handle)
    serfClient.Stop(handle)
    */
  }
}
