package avlib


type AVStream interface {
	LoadAll() []string
	Load(string) (AVIO, bool)
	Store(string, AVIO)
	Del(string)
}
