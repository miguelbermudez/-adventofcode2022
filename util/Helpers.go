package util

import "golang.org/x/exp/constraints"

type Number = interface {
	constraints.Integer | constraints.Float
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

// https://gosamples.dev/generics-reduce/
func Reduce[T, M any](ts []T, fn func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range ts {
		acc = fn(acc, v)
	}
	return acc
}

// https://stackoverflow.com/a/70370013
func SumSlice[T Number](ts []T) T {
	return Reduce(ts, func(acc T, current T) T {
		return acc + current
	}, 0)
}

// https://stackoverflow.com/a/71910002
func Last[E any](s []E) (E, bool) {
	if len(s) == 0 {
		var zero E
		return zero, false
	}
	return s[len(s)-1], true
}
