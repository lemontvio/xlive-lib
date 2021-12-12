package avlib

type AVServer interface {
	Initialize(Config)
	Serve()
	Start()
	Stop()
	Error() chan error
	String() string
	Version() string
}

type AVServerCall func()  AVServer