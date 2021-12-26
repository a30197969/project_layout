package server

import (
	"context"
	"log"
	"net/http"
	"project_layout/internal/conf"
)

type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}

type MyServer struct {
	addr    string
	handler http.Handler
	hs      *http.Server
}

func NewMyServer(cfg *conf.Config) Server {
	return &MyServer{
		addr:    cfg.Http.Addr,
		handler: http.DefaultServeMux,
	}
}

func (m *MyServer) Start(ctx context.Context) error {
	// 创建http.Server对象定制参数
	m.hs = &http.Server{
		Addr:    m.addr,
		Handler: m.handler,
	}
	//go func() {
	//	<-ctx.Done()
	//	log.Printf("server %s stop\n", m.addr)
	//	m.hs.Shutdown(ctx)
	//}()
	log.Printf("http server %s start\n", m.addr)
	return m.hs.ListenAndServe()
}
func (m *MyServer) Stop(ctx context.Context) error {
	<-ctx.Done()
	log.Printf("http server %s stop\n", m.addr)
	return m.hs.Shutdown(ctx)
}
