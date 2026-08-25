package http

import (
	"fmt"
	"net/http"
)

func NewMux(routes Routes) Handler {
	mux := http.NewServeMux()
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		mux.Handle(pattern, With(route.Handler, route.Middleware))
	}
	return mux
}
