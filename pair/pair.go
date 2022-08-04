package pair

type Pair[T, U any] struct {
	First  T
	Second U
}

func First[T, U any](p Pair[T, U]) T {
	return p.First
}

func Second[T, U any](p Pair[T, U]) U {
	return p.Second
}

func MkPair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{first, second}
}

func Map[T, TNew, U, UNew any](f func(T, U) (TNew, UNew), p Pair[T, U]) Pair[TNew, UNew] {
	first, second := f(p.First, p.Second)
	return MkPair(first, second)
}

func (p Pair[T, U]) Do(f func(T, U)) {
	f(p.First, p.Second)
}

func Apply[T, U, V any](f func(T, U) V, p Pair[T, U]) V {
	return f(p.First, p.Second)
}

func MapLeft[T, TNew, U any](f func(T) TNew, p Pair[T, U]) Pair[TNew, U] {
	return MkPair(f(p.First), p.Second)
}

func MapRight[T, U, UNew any](g func(U) UNew, p Pair[T, U]) Pair[T, UNew] {
	return MkPair(p.First, g(p.Second))
}

func (p Pair[T, U]) Unwrap() (T, U) {
	return p.First, p.Second
}

func (p Pair[T, U]) Swap() Pair[U, T] {
	return MkPair(p.Second, p.First)
}
