package gotools

import (
	"sync"
)

func ParallelForEach[T any](
	arr []T,
	delegate func(row T, i int),
	parallelism int,
) {
	wg := sync.WaitGroup{}

	for i := 0; i < len(arr); i++ {
		if i > 0 && i%parallelism == 0 {
			wg.Wait()
		}

		wg.Add(1)
		go func(i int) {
			delegate(arr[i], i)
			wg.Done()
		}(i)
	}
}
