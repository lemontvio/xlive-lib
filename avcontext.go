package avlib

type AVContext interface {
	IsAlive() bool
	Time() int64
	Do()
	Close()
}
