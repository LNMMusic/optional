package optional

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestOption_Some tests the Some constructor.
func TestOption_Some(t *testing.T) {
	t.Run("Some - string", func(t *testing.T) {
		// arrange
		var input string = "hello"
		var output = Option[string]{value: &input}

		// act
		result := Some(input)

		// assert
		require.Equal(t, output, result)
		require.True(t, result.IsSome())
		v, err := result.Unwrap()
		require.NoError(t, err)
		require.Equal(t, input, v)
	})

	t.Run("Some - int", func(t *testing.T) {
		// arrange
		var input int = 1
		var output = Option[int]{value: &input}

		// act
		result := Some(input)

		// assert
		require.Equal(t, output, result)
		require.True(t, result.IsSome())
		v, err := result.Unwrap()
		require.NoError(t, err)
		require.Equal(t, input, v)
	})

	t.Run("Some - float64", func(t *testing.T) {
		// arrange
		var input float64 = 1.0
		var output = Option[float64]{value: &input}

		// act
		result := Some(input)

		// assert
		require.Equal(t, output, result)
		require.True(t, result.IsSome())
		v, err := result.Unwrap()
		require.NoError(t, err)
		require.Equal(t, input, v)
	})

	t.Run("Some - bool", func(t *testing.T) {
		// arrange
		var input bool = true
		var output = Option[bool]{value: &input}

		// act
		result := Some(input)

		// assert
		require.Equal(t, output, result)
		require.True(t, result.IsSome())
		v, err := result.Unwrap()
		require.NoError(t, err)
		require.Equal(t, input, v)
	})

	t.Run("Some - []string", func(t *testing.T) {
		// arrange
		var input []string = []string{"hello", "world"}
		var output = Option[[]string]{value: &input}

		// act
		result := Some(input)

		// assert
		require.Equal(t, output, result)
		require.True(t, result.IsSome())
		v, err := result.Unwrap()
		require.NoError(t, err)
		require.Equal(t, input, v)
	})

	t.Run("Some - []int", func(t *testing.T) {
		// arrange
		var input []int = []int{1, 2, 3}
		var output = Option[[]int]{value: &input}

		// act
		result := Some(input)

		// assert
		require.Equal(t, output, result)
		require.True(t, result.IsSome())
		v, err := result.Unwrap()
		require.NoError(t, err)
		require.Equal(t, input, v)
	})

	t.Run("Some - []float64", func(t *testing.T) {
		// arrange
		var input []float64 = []float64{1.0, 2.0, 3.0}
		var output = Option[[]float64]{value: &input}

		// act
		result := Some(input)

		// assert
		require.Equal(t, output, result)
		require.True(t, result.IsSome())
		v, err := result.Unwrap()
		require.NoError(t, err)
		require.Equal(t, input, v)
	})

	t.Run("Some - []bool", func(t *testing.T) {
		// arrange
		var input []bool = []bool{true, false}
		var output = Option[[]bool]{value: &input}

		// act
		result := Some(input)

		// assert
		require.Equal(t, output, result)
		require.True(t, result.IsSome())
		v, err := result.Unwrap()
		require.NoError(t, err)
		require.Equal(t, input, v)
	})

	t.Run("Some - struct", func(t *testing.T) {
		// arrange
		type testStruct struct {
			value string
		}
		var input testStruct = testStruct{value: "hello"}
		var output = Option[testStruct]{value: &input}

		// act
		result := Some(input)

		// assert
		require.Equal(t, output, result)
		require.True(t, result.IsSome())
		v, err := result.Unwrap()
		require.NoError(t, err)
		require.Equal(t, input, v)
	})
}

// TestOption_None tests the None constructor.
func TestOption_None(t *testing.T) {
	t.Run("None - string", func(t *testing.T) {
		// arrange
		var output = Option[string]{value: nil}

		// act
		result := None[string]()

		// assert
		require.Equal(t, output, result)
		require.False(t, result.IsSome())
		v, err := result.Unwrap()
		require.Equal(t, "", v)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrUnwrapNone)
	})

	t.Run("None - int", func(t *testing.T) {
		// arrange
		var output = Option[int]{value: nil}

		// act
		result := None[int]()

		// assert
		require.Equal(t, output, result)
		require.False(t, result.IsSome())
		v, err := result.Unwrap()
		require.Equal(t, 0, v)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrUnwrapNone)
	})

	t.Run("None - float64", func(t *testing.T) {
		// arrange
		var output = Option[float64]{value: nil}

		// act
		result := None[float64]()

		// assert
		require.Equal(t, output, result)
		require.False(t, result.IsSome())
		v, err := result.Unwrap()
		require.Equal(t, 0.0, v)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrUnwrapNone)
	})

	t.Run("None - bool", func(t *testing.T) {
		// arrange
		var output = Option[bool]{value: nil}

		// act
		result := None[bool]()

		// assert
		require.Equal(t, output, result)
		require.False(t, result.IsSome())
		v, err := result.Unwrap()
		require.Equal(t, false, v)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrUnwrapNone)
	})

	t.Run("None - []string", func(t *testing.T) {
		// arrange
		var output = Option[[]string]{value: nil}

		// act
		result := None[[]string]()

		// assert
		require.Equal(t, output, result)
		require.False(t, result.IsSome())
		v, err := result.Unwrap()
		require.Equal(t, []string(nil), v)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrUnwrapNone)
	})

	t.Run("None - []int", func(t *testing.T) {
		// arrange
		var output = Option[[]int]{value: nil}

		// act
		result := None[[]int]()

		// assert
		require.Equal(t, output, result)
		require.False(t, result.IsSome())
		v, err := result.Unwrap()
		require.Equal(t, []int(nil), v)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrUnwrapNone)
	})

	t.Run("None - []float64", func(t *testing.T) {
		// arrange
		var output = Option[[]float64]{value: nil}

		// act
		result := None[[]float64]()

		// assert
		require.Equal(t, output, result)
		require.False(t, result.IsSome())
		v, err := result.Unwrap()
		require.Equal(t, []float64(nil), v)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrUnwrapNone)
	})

	t.Run("None - []bool", func(t *testing.T) {
		// arrange
		var output = Option[[]bool]{value: nil}

		// act
		result := None[[]bool]()

		// assert
		require.Equal(t, output, result)
		require.False(t, result.IsSome())
		v, err := result.Unwrap()
		require.Equal(t, []bool(nil), v)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrUnwrapNone)
	})

	t.Run("None - struct", func(t *testing.T) {
		// arrange
		type testStruct struct {}
		var output = Option[testStruct]{value: nil}

		// act
		result := None[testStruct]()

		// assert
		require.Equal(t, output, result)
		require.False(t, result.IsSome())
		v, err := result.Unwrap()
		require.Equal(t, testStruct{}, v)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrUnwrapNone)
	})
}