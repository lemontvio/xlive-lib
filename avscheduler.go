package avutil

type AVScheduler interface {
	Call(Config)
	Del(string)
}
