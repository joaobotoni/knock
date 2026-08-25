package http

type Route struct {
	Method     string
	Path       string
	Handler    Handler
	Middleware Stack
}

type Routes []Route

func Router(prefix string, routes Routes) Routes {
	return mount(prefix, routes)
}

func mount(prefix string, routes Routes) Routes {
	mounted := make(Routes, len(routes))
	for i, route := range routes {
		route.Path = prefix + route.Path
		mounted[i] = route
	}
	return mounted
}
