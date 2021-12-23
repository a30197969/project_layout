package pkg

import (
	"context"
	"os"
	"time"
)

type Option func(o *options)

type options struct {
	ctx         context.Context
	sigs        []os.Signal
	servers     []Server
	stopTimeout time.Duration
}

func SetServer(srv ...Server) Option {
	return func(o *options) { o.servers = srv }
}
