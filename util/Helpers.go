package util

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func Reduce[T, M any](ts []T, fn func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range ts {
		acc = fn(acc, v)
	}
	return acc
}

func Last[E any](s []E) (E, bool) {
	if len(s) == 0 {
		var zero E
		return zero, false
	}
	return s[len(s)-1], true
}
