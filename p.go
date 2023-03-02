package p

func T[T any](t T) *T {
	return &t
}
