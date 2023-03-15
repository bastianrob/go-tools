package gotools

func ForEach[T any](arr []T, delegate func(row T, i int)) {
	for i := 0; i < len(arr); i++ {
		delegate(arr[i], i)
	}
}

func Map[T, R any](arr []T, delegate func(row T, i int) R) []R {
	res := make([]R, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = delegate(arr[i], i)
	}

	return res
}

func Reduce[T, R any](arr []T, delegate func(R, T, int) R, initial R) R {
	res := initial
	for i := 0; i < len(arr); i++ {
		res = delegate(res, arr[i], i)
	}

	return res
}

func Find[T any](arr []T, delegate func(row T) bool) (T, int) {
	for i := 0; i < len(arr); i++ {
		if delegate(arr[i]) {
			return arr[i], i
		}
	}

	return *new(T), -1
}

func Filter[T any](arr []T, delegate func(row T) bool) []T {
	res := []T{}
	for i := 0; i < len(arr); i++ {
		if delegate(arr[i]) {
			res = append(res, arr[i])
		}
	}

	return res
}

func Any[T any](arr []T, delegate func(row T) bool) bool {
	for i := 0; i < len(arr); i++ {
		if delegate(arr[i]) {
			return true
		}
	}

	return false
}

func All[T any](arr []T, delegate func(row T) bool) bool {
	for i := 0; i < len(arr); i++ {
		if !delegate(arr[i]) {
			return false
		}
	}

	return true
}

func Chunk[T any](arr []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(arr) {
		arr, chunks = arr[chunkSize:], append(chunks, arr[0:chunkSize:chunkSize])
	}

	return append(chunks, arr)
}
