package router

import "strings"

type dynamicPathValues = map[string]string

type PathNode struct {
	value     string
	isDynamic bool
	children  []*PathNode
	handlers  []HandleFunc
	method    string
}

func findHandlers(node *PathNode, path, method string) ([]HandleFunc, dynamicPathValues) {
	path = strings.TrimPrefix(path, "/")
	if path == "" {
		return nil, nil
	}
	segments := strings.Split(path, "/")

	handlers, values, _ := traverse(node, segments, 0, make(dynamicPathValues), method)

	return handlers, values
}

func traverse(node *PathNode, segments []string, cursor int, values dynamicPathValues, method string) ([]HandleFunc, dynamicPathValues, bool) {
	if cursor >= len(segments) {
		return nil, nil, false
	}

	if !node.isDynamic && segments[cursor] != node.value {
		return nil, nil, false
	}

	if node.isDynamic {
		key := strings.TrimFunc(node.value, func(r rune) bool {
			return r == '{' || r == '}'
		})
		values = copyValues(values)
		values[key] = segments[cursor]
	}

	if cursor == len(segments)-1 {
		var handlers []HandleFunc

		if node.method == method || node.method == "" {
			handlers = append(handlers, node.handlers...)
		}

		for _, childNode := range node.children {
			if childNode.value == "" && (childNode.method == method || childNode.method == "") {
				handlers = append(handlers, childNode.handlers...)
			}
		}

		if len(handlers) > 0 {
			return handlers, values, true
		}
		return nil, nil, false
	}

	for _, childNode := range node.children {
		if childNode.method == method || childNode.method == "" {
			if handlers, vals, ok := traverse(childNode, segments, cursor+1, values, method); ok {
				return handlers, vals, true
			}
		}
	}

	return nil, nil, false
}

func copyValues(original dynamicPathValues) dynamicPathValues {
	newMap := make(dynamicPathValues)
	for k, v := range original {
		newMap[k] = v
	}
	return newMap
}
