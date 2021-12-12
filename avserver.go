package xlive_lib

type AVServer interface {
	Serve()
	Start()
	Stop()
	Error() chan error
}
