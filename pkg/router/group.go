package router

type Group struct {
}

func (g *Group) Group(path string, middlewares ...[]HandleFunc) *Group {
	// TODO
	return nil
}

func (g *Group) POST(path string, handler HandleFunc) {

}
