package internal

type Result[T any] struct {
	Value T
	Err   error
	Ok    bool
}

func Success[T any](value T) Result[T] {
	return Result[T]{Value: value, Ok: true}
}

func Failure[T any](err error) Result[T] {
	return Result[T]{Err: err}
}

func (r Result[T]) IsOk() bool {
	return r.Ok
}

func (r Result[T]) Unwrap() T {
	if !r.Ok {
		panic(r.Err)
	}
	return r.Value
}

func (r Result[T]) UnwrapOr(def T) T {
	if !r.Ok {
		return def
	}
	return r.Value
}
