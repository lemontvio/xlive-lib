package xlive_lib

type AVScheduler interface {
	Call(Config)
	Del(string)
}
