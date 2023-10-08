package optional

import (
	"encoding/json"
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
func (o *Option[T]) Unwrap() (t T, err error) {
	if o.value == nil {
		err = ErrUnwrapNone
		return
	}
	t = *o.value
	return
}

// UnmarshalJSON indicates how to unmarshal a json value into an Option.
// unmarshalling (always work with reference)
// json.Unmarshal: always sends a ptr of the field
// - interface: always accepts ptr | no matter method receiver
// 
// cases
// - method [ptr-receiver] -> works
// - method [non-ptr]      -> does not work (can't access to the field of &o.Value)
func (o *Option[T]) UnmarshalJSON(data []byte) (err error) {
	err = json.Unmarshal(data, &o.value)
	return
}

// MarshalJSON indicates how to marshal an Option into a json value.
// marshalling (always work with value)
// json.Marshal: always sends a non-ptr of the field
// - interface: not always accepts non-ptr
//              > case non-ptr: works
//              > case ptr: does not work (go rules)
//
// cases
// - method [non-ptr] -> works ()
func (o Option[T]) MarshalJSON() (data []byte, err error) {
	data, err = json.Marshal(o.value)
	return
}