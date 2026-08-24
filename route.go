package knock

import (
	"net/http"
)

type Handler http.Handler

type Middleware func(Handler) Handler

type Stack []Middleware

func (s Stack) Then(h Handler) Handler {
	for i := len(s) - 1; i >= 0; i-- {
		h = s[i](h)
	}
	return h
}

type Route struct {
	Method     string
	Path       string
	Handler    Handler
	Middleware []Middleware
}

type Router interface {
	Routes() []Route
}
