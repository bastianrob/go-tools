package gotools

import (
	"sync"
)

type Runnable[T any] struct {
	workerSize int
	channel    chan T
	delegate   Delegate[T]
	wg         *sync.WaitGroup
}

type Delegate[T any] func(id int, this *Runnable[T], message T)

func NewRunnable[T any](
	workerSize int,
	delegate Delegate[T],
) *Runnable[T] {
	return &Runnable[T]{
		workerSize: workerSize,
		delegate:   delegate,
		channel:    make(chan T, workerSize),
		wg:         new(sync.WaitGroup),
	}
}

func (r *Runnable[T]) Publish(messsage T) {
	r.wg.Add(1)
	r.channel <- messsage
}

func (r *Runnable[T]) Close() {
	close(r.channel)
}

func (r *Runnable[T]) Run() {
	for i := 1; i <= r.workerSize; i++ {
		go func(workerId int) {
			for m := range r.channel {
				r.delegate(workerId, r, m)
				r.wg.Done()
			}
		}(i)
	}
}

func (r *Runnable[T]) Wait() {
	r.wg.Wait()
}
