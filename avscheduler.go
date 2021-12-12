package avlib

type AVScheduler interface {
	Call(Config)
	Del(string)
}
