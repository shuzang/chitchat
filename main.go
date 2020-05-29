package main

import (
	"log"
	"net/http"

	. "github.com/shuzang/chitchat/routes"
)

func main() {
	startWebServer("8080")
}

// 通过指定端口启动服务器
func startWebServer(port string) {
	r := NewRouter()

	assets := http.FileServer(http.Dir("public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r) // 通过 router.go 中定义的路由器进行分发

	log.Println("Starting HTTP server at " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
