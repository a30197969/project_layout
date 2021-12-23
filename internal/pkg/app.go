package pkg

import (
	"context"
	"golang.org/x/sync/errgroup"
	"sync"
)

type App struct {
	opts   options
	ctx    context.Context
	cancel func()
}

func NewApp(opts ...Option) *App {
	// 配置文件设置
	o := options{
		ctx: context.Background(),
	}
	for _, opt := range opts {
		opt(&o)
	}
	ctx, cancel := context.WithCancel(o.ctx)
	return &App{
		ctx:    ctx,
		cancel: cancel,
		opts:   o,
	}
}

func (a *App) Run() error {
	eg, ctx := errgroup.WithContext(a.ctx)
	wg := sync.WaitGroup{}
	for _, srv := range a.opts.servers {
		srv := srv
		eg.Go(func() error {
			<-ctx.Done() // wait for stop signal
			sctx, cancel := context.WithTimeout(context.Background(), a.opts.stopTimeout)
			defer cancel()
			return srv.Stop(sctx)
		})
		wg.Add(1)
		eg.Go(func() error {
			wg.Done()
			return srv.Start(ctx)
		})
	}
	wg.Wait()
	return nil
}
func (a *App) Stop() error {
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}
