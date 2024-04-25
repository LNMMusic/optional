package nullable

// Null is a generic type that represents a nullable value.
// There are 2 possible states:
// - Null:
//   > value: zero value of T
//   > valid: false
//
// - Not Null:
//   > value: value of T
//   > valid: true
type Null[T any] struct {
	value T
	valid bool
}

// Constructors
// None returns a Null[T] with valid set to false.
func None[T any]() Null[T] {
	return Null[T]{
		valid: false,
	}
}

// Some returns a Null[T] with valid set to true.
func Some[T any](value T) Null[T] {
	return Null[T]{
		value: value,
		valid: true,
	}
}

// Methods
// - Inspection
// IsNull returns true if the Null is in a Null state.
func (n Null[T]) IsNull() bool {
	return !n.valid
}

// - Fetching
// Unwrap returns the inner value of a Not Null.
// If the Null is in a Null state, Unwrap panics.
func (n Null[T]) Unwrap() T {
	if n.valid {
		return n.value
	}
	panic("unable to unwrap Null")
}
