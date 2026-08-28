package khttp

import "slices"

type Middleware func(Handler) Handler

type Stack []Middleware

func With(next Handler, stack Stack) Handler {
	handler := next
	for _, middleware := range slices.Backward(stack) {
		handler = middleware(handler)
	}
	return handler
}
