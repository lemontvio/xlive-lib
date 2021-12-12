package xlive_lib

type AVFormat interface {
	Initialize(Config)
	Version() string
	String() string
	Do()
	Close()
	Error() chan error
}

type AVFormatCall func() AVFormat
