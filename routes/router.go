package routes

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	// 创建 mux.Router 路由实例
	router := mux.NewRouter().StrictSlash(true)

	// 遍历 routes.go 中所有 webRoutes
	for _, route := range webRoutes {
		//将每个路由应用到路由器
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}
