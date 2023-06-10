package util

import (
	"golang.org/x/exp/constraints"
)

func RemoveFromSlice[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func RemovePointerFromSlice[T any](s []*T, i int) []*T {
	s[i] = s[len(s)-1]
	s[len(s)-1] = nil
	return s[:len(s)-1]
}

func OrderedRemoveFromSlice[T any](s []T, i int) []T {
	return append(s[:i], s[i+1:]...)
}

func OrderedRemovePointerFromSlice[T any](s []*T, i int) []*T {
	s[i] = nil
	return append(s[:i], s[i+1:]...)
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}

	return b
}
