package optional

import (
	"encoding/json"
	"errors"
)

var (
	ErrUnwrapNone = errors.New("cannot unwrap None")
)

// Option is a type that represents an optional value.
type Option[T any] struct {
	Value *T
}

// methods
// -> is some: true if the option is a Some value
func (o *Option[T]) IsSome() bool {
	return o.Value != nil
}
// -> Unwrap: returns the inner T of a Some. Panics if the self value equals None.
func (o *Option[T]) Unwrap() (t T, err error) {
	if o.Value == nil {
		err = ErrUnwrapNone
		return
	}
	t = *o.Value
	return
}

// constructors
func Some[T any](Value T) Option[T] {
	return Option[T]{Value: &Value}
}
func None[T any]() Option[T] {
	return Option[T]{Value: nil}
}


// unmarshalling (always work with reference)
// json.Unmarshal: always sends a ptr of the field
// - interface: always accepts ptr | no matter method receiver
// 
// cases
// - method [ptr-receiver] -> works
// - method [non-ptr]      -> does not work (can't access to the field of &o.Value)
func (o *Option[T]) UnmarshalJSON(data []byte) (err error) {
	err = json.Unmarshal(data, &o.Value)
	return
}

// marshalling (always work with value)
// json.Marshal: always sends a non-ptr of the field
// - interface: not always accepts non-ptr
//              > case non-ptr: works
//              > case ptr: does not work (go rules)
//
// cases
// - method [non-ptr] -> works ()
func (o Option[T]) MarshalJSON() (data []byte, err error) {
	data, err = json.Marshal(o.Value)
	return
}