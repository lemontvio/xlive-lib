package avutil

import (
	"context"

	"github.com/lemontvio/go-requests"
)

type Config struct {
	Addr             string
	Proxy            string
	Root             string
	Ptl              string
	ID               string
	URL              string
	HTTPCrawlerProxy string
	HTTPProxy        string
	Def              int
	Size             int
	Duration         float64
	Params           map[string]interface{}
	Options          *requests.Options
	Requests         *requests.Requests
	Callabck         func(context.Context)
	Streams          AVStream
	IO               AVIO
	Scheduler        AVScheduler
	Logger           AVLogger
}
