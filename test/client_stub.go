package test

type ClientTest struct {
  UserEvent UserEventCall
  UserEventReturn error

  ListKeys bool
  ListKeysRet ListKeysReturn

  Stats bool
  StatsRet StatsReturn

  UseKey string
  UseKeyRet UseKeyReturn

  Leave bool
  LeaveRet LeaveReturn

  ForceLeave string
  ForceLeaveRet error

  Join JoinCall
  JoinRet JoinReturn

  Members bool
  MembersRet MembersReturn
}

func (c *ClientTest) UserEvent(name string, payload []byte, coalesce bool) error {
  c.UserEvent = UserEventCall{Name: name, Payload: payload, Coalesce: coalesce}
  return c.UserEventReturn
}

func (c *ClientTest) ListKeys() (map[string]int, int, map[string]string, error) {
  c.ListKeys = true
  ret := c.ListKeysReturn
  return ret.First, ret.Second, ret.Third, ret.Fourth
}

func (c *ClientTest) Stats() (map[string]map[string]string, error) {
  c.Stats = true
  ret := c.StatsReturn
  return ret.First, ret.Second
}
func (c *ClientTest) UseKey(key string) (map[string]string, error) {

}

func (c *ClientTest) Leave() error{

}

func (c *ClientTest) ForceLeave(node string) error {

}

func (c *ClientTest) Join(addrs []string, replay bool) (int, error) {

}

func (c *ClientTest) Members() ([]SerfClient.Member, error) {

}
