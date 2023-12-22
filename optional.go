package optional

import (
	"errors"
)

var (
	ErrUnwrapNone = errors.New("cannot unwrap None")
)

// Constructors
// Some returns an Option with a Some value.
func Some[T any](value T) Option[T] {
	return Option[T]{value: &value}
}
// None returns an Option with a None value.
func None[T any]() Option[T] {
	return Option[T]{value: nil}
}


// Option is a type that represents an optional value.
// - value is inmutable
// - concurrency safe
type Option[T any] struct {
	value *T
}

// Methods
// IsSome returns true if the option is a Some value.
func (o *Option[T]) IsSome() bool {
	return o.value != nil
}
// Unwrap returns a copy of the inner value of a Some.
// If the Option is a None, Unwrap panics.
func (o *Option[T]) Unwrap() (t T) {
	if o.value == nil {
		panic("unable to unwrap None")
	}
	t = *o.value
	return
}
