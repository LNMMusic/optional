package nullable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNull_Constructor_None(t *testing.T) {
	t.Run("should return a Null[T] of type int with valid set to false", func(t *testing.T) {
		// arrange
		// ...

		// act
		intNullable := None[int]()

		// assert
		require.Equal(t, intNullable.value, 0)
		require.False(t, intNullable.valid)
	})

	t.Run("should return a Null[T] of type string with valid set to false", func(t *testing.T) {
		// arrange
		// ...

		// act
		stringNullable := None[string]()

		// assert
		require.Equal(t, stringNullable.value, "")
		require.False(t, stringNullable.valid)
	})

	t.Run("should return a Null[T] of type bool with valid set to false", func(t *testing.T) {
		// arrange
		// ...

		// act
		boolNullable := None[bool]()

		// assert
		require.Equal(t, boolNullable.value, false)
		require.False(t, boolNullable.valid)
	})

	t.Run("should return a Null[T] of type float64 with valid set to false", func(t *testing.T) {
		// arrange
		// ...

		// act
		float64Nullable := None[float64]()

		// assert
		require.Equal(t, float64Nullable.value, 0.0)
		require.False(t, float64Nullable.valid)
	})

	t.Run("should return a Null[T] of type struct with valid set to false", func(t *testing.T) {
		// arrange
		type Person struct {
			Name string
		}

		// act
		structNullable := None[Person]()

		// assert
		require.Equal(t, structNullable.value, Person{})
		require.False(t, structNullable.valid)
	})

	t.Run("should return a Null[T] of type interface{} with valid set to false", func(t *testing.T) {
		// arrange
		// ...

		// act
		interfaceNullable := None[interface{}]()

		// assert
		require.Nil(t, interfaceNullable.value)
		require.False(t, interfaceNullable.valid)
	})
}

func TestNull_Constructor_Some(t *testing.T) {
	t.Run("should return a Null[T] of type int with valid set to true", func(t *testing.T) {
		// arrange
		// ...

		// act
		intNullable := Some[int](42)

		// assert
		require.Equal(t, intNullable.value, 42)
		require.True(t, intNullable.valid)
	})

	t.Run("should return a Null[T] of type string with valid set to true", func(t *testing.T) {
		// arrange
		// ...

		// act
		stringNullable := Some[string]("hello")

		// assert
		require.Equal(t, stringNullable.value, "hello")
		require.True(t, stringNullable.valid)
	})

	t.Run("should return a Null[T] of type bool with valid set to true", func(t *testing.T) {
		// arrange
		// ...

		// act
		boolNullable := Some[bool](true)

		// assert
		require.Equal(t, boolNullable.value, true)
		require.True(t, boolNullable.valid)
	})

	t.Run("should return a Null[T] of type float64 with valid set to true", func(t *testing.T) {
		// arrange
		// ...

		// act
		float64Nullable := Some[float64](3.14)

		// assert
		require.Equal(t, float64Nullable.value, 3.14)
		require.True(t, float64Nullable.valid)
	})

	t.Run("should return a Null[T] of type struct with valid set to true", func(t *testing.T) {
		// arrange
		type Person struct {
			Name string
		}

		// act
		structNullable := Some[Person](Person{Name: "John"})

		// assert
		require.Equal(t, structNullable.value, Person{Name: "John"})
		require.True(t, structNullable.valid)
	})

	t.Run("should return a Null[T] of type interface{} with valid set to true", func(t *testing.T) {
		// arrange
		// ...

		// act
		interfaceNullable := Some[interface{}](42)

		// assert
		require.Equal(t, interfaceNullable.value, 42)
		require.True(t, interfaceNullable.valid)
	})

	// pointer types
	t.Run("should return a Null[T] of type *int with valid set to true", func(t *testing.T) {
		// arrange
		// ...

		// act
		intPtr := 42
		intPtrNullable := Some[*int](&intPtr)

		// assert
		require.Equal(t, *intPtrNullable.value, 42)
		require.True(t, intPtrNullable.valid)
	})
}

func TestNull_IsNull(t *testing.T) {
	t.Run("should return true if the Null is in a Null state", func(t *testing.T) {
		// arrange
		intNullable := None[int]()

		// act
		isNull := intNullable.IsNull()

		// assert
		require.True(t, isNull)
	})

	t.Run("should return false if the Null is in a Not Null state", func(t *testing.T) {
		// arrange
		intNullable := Some[int](42)

		// act
		isNull := intNullable.IsNull()

		// assert
		require.False(t, isNull)
	})
}

func TestNull_Unwrap(t *testing.T) {
	t.Run("should return the inner value of a Not Null", func(t *testing.T) {
		// arrange
		intNullable := Some[int](42)

		// act
		value := intNullable.Unwrap()

		// assert
		require.Equal(t, value, 42)
	})

	t.Run("should panic if the Null is in a Null state", func(t *testing.T) {
		// arrange
		intNullable := None[int]()

		// act
		require.PanicsWithValue(t, "unable to unwrap Null", func() {
			intNullable.Unwrap()
		})
	})
}
