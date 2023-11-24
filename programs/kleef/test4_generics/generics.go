package main

func id[T any](in T) T {
	return in
}
