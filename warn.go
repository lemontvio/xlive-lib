package avlib

type Warn interface {
	Initialize(map[string]interface{}) error
	WriteError(string)
	WriteStream(string, bool)
	String()string
	Close() error
}
