package main

import (
	"fmt"
	"net/http"
	"project_layout/internal/conf"
	"project_layout/internal/pkg"
	"project_layout/internal/server"
)

func newApp(hs server.Server) *pkg.App {
	return pkg.NewApp(pkg.SetServer(hs))
}
func Thread(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Thread!")
}
func main() {
	// 获取配置文件
	confPath := "configs/main.yaml"
	conf := conf.GetConf(confPath)
	fmt.Printf("%+v\n", conf)
	app := initApp(conf)
	// 应用启动
	if err := app.Run(); err != nil {
		panic(err)
	}
}
