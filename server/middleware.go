package server

type Middleware func(Handler) Handler

type Stack []Middleware

func With(next Handler, stack Stack) Handler {
	handler := next
	for i := len(stack) - 1; i >= 0; i-- {
		handler = stack[i](handler)
	}
	return handler
}
