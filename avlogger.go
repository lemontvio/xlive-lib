package avlib

import "io"

type AVLogger interface {
	Trak(...interface{})
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})

	Trakf(string, ...interface{})
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})

	SetLevel(int)
	SetLevel2(string)
	SetWriter(io.Writer)

	New(...string) interface{}
	Close()
	//
	io.Writer
}

func Logger(v interface{}) AVLogger {
	return v.(AVLogger)
}
