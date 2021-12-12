package avlib

type AVServer interface {
	Serve()
	Start()
	Stop()
	Error() chan error
}
