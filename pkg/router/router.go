package router

import "net/http"

type Router struct {
}

func DefaultRouter() *Router {
	r := Router{}
	http.Handle("/", http.HandlerFunc(r.baseHandle))
	return &r
}

func (r *Router) Group(path string, middlewares ...[]HandleFunc) *Group {
	// TODO
	return nil
}

func (r *Router) baseHandle(writer http.ResponseWriter, request *http.Request) {

}
