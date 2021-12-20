package avlib

type AVScheduler interface {
	Call(Config)
	Del(string)
	Streams() map[string]map[bool]int64
}
