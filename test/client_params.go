package test

type UserEventCall struct {
  Name string
  Payload []byte
  Coalesce bool
}

type JoinCall struct {
  Addrs []string
  Replay bool
}

type ListKeysReturn struct {
  First map[string]int
  Second int
  Third map[string]string
  Fourth error
}

type StatsReturn struct {
  First map[string]map[string]string
  Second error
}

type UseKeyReturn struct {
  First map[string]string
  Second error
}
