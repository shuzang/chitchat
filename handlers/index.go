package handlers

import (
	"net/http"

	"github.com/shuzang/chitchat/models"
)

// 论坛首页路由的处理器方法
func Index(writer http.ResponseWriter, request *http.Request) {
	threads, err := models.Threads()
	if err == nil {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, threads, "layout", "navbar", "index")
		} else {
			generateHTML(writer, threads, "layout", "auth.navbar", "index")
		}
	}
}
