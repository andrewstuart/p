package p

import "fmt"

// Return a pointer to anything
func T[T any](t T) *T {
	return &t
}

// Return a pointer to a printf'd string
func F(tpl string, args ...any) *string {
	return T(fmt.Sprintf(tpl, args...))
}
