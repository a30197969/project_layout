package main

import (
	"fmt"
	"net/http"
	"project_layout/internal/conf"
	"project_layout/internal/pkg"
)

func initApp(hs pkg.Server) (*pkg.App, error) {
	confPath := "configs/main.yaml"
	conf := conf.GetConf(confPath)
	fmt.Printf("%+v\n", conf)
	//c := GetRedisConf()
	//r := NewRedis(c)
	//greeterRepo := data.NewGreeterRepo(dataData, logger)
	app := newApp(hs)
	return app, nil
}

func newApp(hs pkg.Server) *pkg.App {
	return pkg.NewApp(pkg.SetServer(hs))
}
func Thread(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Thread!")
}
func main() {
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/thread", Thread)
	hs := pkg.NewMyServer("127.0.0.1:8001", mux1)
	app, err := initApp(hs)
	if err != nil {
		panic(err)
	}
	// 应用启动
	if err := app.Run(); err != nil {
		panic(err)
	}
}
