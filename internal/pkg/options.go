package pkg

import (
	"context"
	"os"
	"project_layout/internal/server"
	"time"
)

type Option func(o *options)

type options struct {
	ctx         context.Context
	sigs        []os.Signal
	servers     []server.Server
	stopTimeout time.Duration
}

func SetServer(srv ...server.Server) Option {
	return func(o *options) { o.servers = srv }
}
