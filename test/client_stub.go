package test

import (
  SerfClient "github.com/hashicorp/serf/client"
)

type ClientTest struct {
  UserEventCalled bool
  ListKeysCalled bool
  StatsCalled bool
  UseKeyCalled bool
  LeaveCalled bool
  ForceLeaveCalled bool
  JoinCalled bool
  MembersCalled bool
}

func NewClientTest() *ClientTest {
  return &ClientTest{}
}

func (c *ClientTest) UserEvent(name string, payload []byte, coalesce bool) error {
  c.UserEventCalled = true
  return nil
}

func (c *ClientTest) ListKeys() (map[string]int, int, map[string]string, error) {
  c.ListKeysCalled = true
  return map[string]int{}, 0, map[string]string{}, nil
}

func (c *ClientTest) Stats() (map[string]map[string]string, error) {
  c.StatsCalled = true
  return map[string]map[string]string{}, nil
}
func (c *ClientTest) UseKey(key string) (map[string]string, error) {
  c.UseKeyCalled = true
  return map[string]string{}, nil
}

func (c *ClientTest) Leave() error {
  c.LeaveCalled = true
  return nil
}

func (c *ClientTest) ForceLeave(node string) error {
  c.ForceLeaveCalled = true
  return nil
}

func (c *ClientTest) Join(addrs []string, replay bool) (int, error) {
  c.JoinCalled = true
  return 0, nil
}

func (c *ClientTest) Members() ([]SerfClient.Member, error) {
  c.MembersCalled = true
  return []SerfClient.Member{}, nil
}
