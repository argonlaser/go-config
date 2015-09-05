/*
Example program.

View help and defaults:
  $ go run _example/main.go -h

Assign some values:
  $ go run _example/main.go --timeout 10s --cache-size 100mb

Assign to the nsqlookup list:
  $ go run _example/main.go --nsq-lookup foo,bar --nsq-lookup baz

Multiple resolvers:
  $ LOG_LEVEL=error go run _example/main.go --concurrency 100

Validation:
  $ go run _example/main.go --batch-size 0
  $ go run _example/main.go --batch-size 1500
  $ go run _example/main.go --concurrency -5
*/
package main

import (
	"log"
	"time"

	"github.com/tj/go-config"
)

type Options struct {
	Timeout     time.Duration `desc:"message timeout"`
	Concurrency uint          `desc:"max in-flight messages"`
	CacheSize   config.Bytes  `desc:"cache size in bytes"`
	BatchSize   uint          `desc:"batch size" validate:"min=1,max=1000"`
	LogLevel    string        `desc:"set the log severity" from:"env,flag"`
	NSQ         struct {
		Address     string   `desc:"nsqd address"`
		Lookup      []string `desc:"nsqlookupd addresses"`
		MaxInFlight int      `desc:"nsqd max in flight messages"`
	}
}

func main() {
	options := &Options{
		Timeout:     time.Second * 5,
		Concurrency: 10,
		CacheSize:   config.ParseBytes("500mb"),
		BatchSize:   1000,
		LogLevel:    "info",
	}

	config.MustResolve(options)
	log.Printf("%+v", options)
}
