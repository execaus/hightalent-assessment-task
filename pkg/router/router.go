package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type Router struct {
	requestTimeout time.Duration
	rootPathNode   *PathNode
	mux            *http.ServeMux
}

const defaultPort = "8080"
const defaultRequestTime = time.Second * 5

func DefaultRouter() *Router {
	r := Router{
		requestTimeout: defaultRequestTime,
	}

	r.mux = http.NewServeMux()

	return &r
}

func (r *Router) GetServer(port *string) *http.Server {
	listenPort := defaultPort

	if port != nil {
		listenPort = *port
	}

	return &http.Server{
		Addr:    ":" + listenPort,
		Handler: r.mux,
	}
}

func (r *Router) Group(path string, middlewares ...HandleFunc) *Group {
	node := PathNode{
		value:    path,
		handlers: middlewares,
	}

	r.rootPathNode = &node

	return &Group{
		node: &node,
	}
}

func (r *Router) BaseHandle(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(request.Context(), r.requestTimeout)

	requestContext := NewRequestContext(ctx, cancel, writer, request)

	handlers, dynamicValues := r.getHandler(request.URL.Path, request.Method)
	if len(handlers) == 0 {
		log.Println("no handler found for path: " + request.URL.Path)
		requestContext.SendNotFound("handler not found for the requested path")
		return
	}

	requestContext.DynamicValues = dynamicValues

	for _, handler := range handlers {
		handler(requestContext)

		if requestContext.isAbort {
			return
		}
	}
}

func (r *Router) getHandler(path, method string) ([]HandleFunc, dynamicPathValues) {
	return findHandlers(r.rootPathNode, path, method)
}

func (r *Router) PrintRoutes() {
	traversePrint(r.rootPathNode, "")
}

func traversePrint(node *PathNode, currentPath string) {
	var printValue, fullPath string

	if node.value == "" {
		printValue = node.value
	} else {
		printValue = node.value + "/"
	}

	if len(node.children) == 0 {
		printPath := strings.TrimSuffix(currentPath+printValue, "/")
		fmt.Printf("%s%s%s [%v handlers]\n", printPath, strings.Repeat("\t", 6), node.method, len(node.handlers))
		return
	} else {
		fullPath = currentPath + printValue
	}

	for _, child := range node.children {
		traversePrint(child, fullPath)
	}
}
