package channelutils

import (
	"channels/results"
	"channels/union"
	"errors"
)

func MapOn[T, U any](mapper func(T) U, cin <-chan T) <-chan U {
	cout := make(chan U)
	go func() {
		for {
			x := <-cin
			cout <- mapper(x)
		}
	}()
	return cout
}

func MapOnResult[T, U any](mapper func(T) U, cin <-chan results.Result[T]) <-chan results.Result[U] {
	resultMapper := func(o results.Result[T]) results.Result[U] {
		return results.MapOk(mapper, o)
	}
	return MapOn(resultMapper, cin)
}

func Transfer[T any](cin <-chan T, cout chan<- T) {
	go func() {
		for {
			x := <-cin
			cout <- x
		}
	}()
}

func FanIn[T, U any](c1 <-chan T, c2 <-chan U) <-chan union.Union[T, U] {
	c := make(chan union.Union[T, U])

	go Transfer(MapOn(union.MkLeft[T, U], c1), c)
	go Transfer(MapOn(union.MkRight[T, U], c2), c)

	return c
}

func Iterate[T any](initial T, f func(T) T) <-chan T {
	c := make(chan T)

	go func() {
		for {
			c <- initial
			initial = f(initial)
		}
	}()

	return c
}

func Take[T any](n uint64, cin <-chan T) <-chan results.Result[T] {
	c := make(chan results.Result[T])

	go func() {
		for i := uint64(0); i < n; i += 1 {
			x := <-cin
			c <- results.Ok(x)
		}
		c <- results.Err[T](errors.New("Iteration Ended"))
	}()

	return c
}

func Drop[T any](n uint64, cin <-chan T) <-chan T {
	go func() {
		for i := uint64(0); i < n; i += 1 {
			<-cin
		}
	}()
	return cin
}
