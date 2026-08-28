package khttp

import (
	"fmt"
	"net/http"
)

func NewMux(routes Routes) Handler {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(pattern(route.Method, route.Path), With(route.Handler, route.Middleware))
	}
	return mux
}

func pattern(method, path string) string {
	return fmt.Sprintf("%s %s", method, path)
}
