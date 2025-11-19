package router

import (
	"net/http"
	"strings"
)

type Group struct {
	node *PathNode
}

func (g *Group) Group(path string, middlewares ...HandleFunc) *Group {
	value := strings.Trim(path, "/")
	isDynamic := strings.Contains(path, "{")

	group := Group{
		node: &PathNode{
			value:     value,
			isDynamic: isDynamic,
			handlers:  middlewares,
		},
	}

	g.node.children = append(g.node.children, group.node)

	return &group
}

func (g *Group) POST(path string, handlers ...HandleFunc) {
	value := strings.Trim(path, "/")
	isDynamic := strings.Contains(path, "{")

	node := PathNode{
		value:     value,
		isDynamic: isDynamic,
		method:    http.MethodPost,
		handlers:  handlers,
	}

	g.node.children = append(g.node.children, &node)
}
