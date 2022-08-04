package union

import "errors"

type Union[T, U any] struct {
	left   T
	right  U
	isLeft bool
}

func MkLeft[T, U any](left T) Union[T, U] {
	return Union[T, U]{left: left, isLeft: true}
}

func MkRight[T, U any](right U) Union[T, U] {
	return Union[T, U]{right: right, isLeft: false}
}

func (u *Union[T, U]) AccessLeft() (*T, error) {
	if !u.isLeft {
		return nil, errors.New("Cannot access as Left as it contains Right value!")
	}
	return &u.left, nil
}

func (u *Union[T, U]) AccessRight() (*U, error) {
	if u.isLeft {
		return nil, errors.New("Cannot access as Right as it contains Left value!")
	}
	return &u.right, nil
}

func Map[T, TNew, U, UNew any](f func(T) TNew, g func(U) UNew, u Union[T, U]) Union[TNew, UNew] {
	a, _ := MapLeft(f, u)
	ret, _ := MapRight(g, *a)
	return *ret
}

func (u Union[T, U]) Do(f func(T), g func(U)) {
	if u.isLeft {
		f(u.left)
	} else {
		g(u.right)
	}
}

func Apply[T, U, V any](f func(T) V, g func(U) V, u Union[T, U]) V {
	if u.isLeft {
		return f(u.left)
	} else {
		return g(u.right)
	}
}

func MapLeft[T, TNew, U any](f func(T) TNew, u Union[T, U]) (*Union[TNew, U], error) {
	left, err := u.AccessLeft()
	if err != nil {
		return nil, err
	}
	ret := MkLeft[TNew, U](f(*left))
	return &ret, nil
}

func MapRight[T, U, UNew any](f func(U) UNew, u Union[T, U]) (*Union[T, UNew], error) {
	right, err := u.AccessRight()
	if err != nil {
		return nil, err
	}
	ret := MkRight[T](f(*right))
	return &ret, nil
}
