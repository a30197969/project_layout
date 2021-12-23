package pkg

import (
	"context"
	"log"
	"net/http"
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

func NewMyServer(addr string, handler http.Handler) Server {
	return &MyServer{
		addr:    addr,
		handler: handler,
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
