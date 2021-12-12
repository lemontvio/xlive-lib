package avutil

type AVCrawler interface {
	Initialize(config Config)
	String() string
	Version() string
	Suitable(string) bool
	Do() (Config, error)
}

type AVCrawlerCall func() AVCrawler
