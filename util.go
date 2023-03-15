package gotools

func If[T any](condition bool, truthy, falsy T) T {
	if condition {
		return truthy
	}

	return falsy
}
