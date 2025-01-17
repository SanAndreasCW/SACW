package commons

func If[T any](condition bool, a T, b T) T {
	if condition {
		return T(a)
	} else {
		return T(b)
	}
}
