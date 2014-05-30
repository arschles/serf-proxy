package main

import SerfClient "github.com/hashicorp/serf/client"

//this is a interface that github.com/hashicorp/serf/client adheres to.
//we use it here instead of the serf client directly so we can "inject" it
type Client interface {
  UserEvent(name string, payload []byte, coalesce bool) error
  ListKeys() (map[string]int, int, map[string]string, error)
  Stats() (map[string]map[string]string, error)
  UseKey(key string) (map[string]string, error)
  Leave() error
  ForceLeave(node string) error
  Join(addrs []string, replay bool) (int, error)
  Members() ([]SerfClient.Member, error)
}
