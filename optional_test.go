package optional

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestOption_Some tests the Some constructor.
func TestOption_Some(t *testing.T) {
	t.Run("Some - string", func(t *testing.T) {
		// arrange
		// ...
		
		// act
		input := "hello"
		result := Some(input)

		// assert
		// - optional equality
		expected := Option[string]{value: new(string)}; *expected.value = "hello"
		require.Equal(t, expected, result)
		// - check if it is a Some
		require.True(t, result.IsSome())
		// - check unwrapped value
		expectedUnwrap := "hello"
		require.Equal(t, expectedUnwrap, result.Unwrap())
	})

	t.Run("Some - int", func(t *testing.T) {
		// arrange
		// ...
		
		// act
		input := 42
		result := Some(input)

		// assert
		// - optional equality
		expected := Option[int]{value: new(int)}; *expected.value = 42
		require.Equal(t, expected, result)
		// - check if it is a Some
		require.True(t, result.IsSome())
		// - check unwrapped value
		expectedUnwrap := 42
		require.Equal(t, expectedUnwrap, result.Unwrap())
	})

	t.Run("Some - float64", func(t *testing.T) {
		// arrange
		// ...
		
		// act
		input := 42.0
		result := Some(input)

		// assert
		// - optional equality
		expected := Option[float64]{value: new(float64)}; *expected.value = 42.0
		require.Equal(t, expected, result)
		// - check if it is a Some
		require.True(t, result.IsSome())
		// - check unwrapped value
		expectedUnwrap := 42.0
		require.Equal(t, expectedUnwrap, result.Unwrap())
	})

	t.Run("Some - bool", func(t *testing.T) {
		// arrange
		// ...
		
		// act
		input := true
		result := Some(input)

		// assert
		// - optional equality
		expected := Option[bool]{value: new(bool)}; *expected.value = true
		require.Equal(t, expected, result)
		// - check if it is a Some
		require.True(t, result.IsSome())
		// - check unwrapped value
		expectedUnwrap := true
		require.Equal(t, expectedUnwrap, result.Unwrap())
	})

	t.Run("Some - []string", func(t *testing.T) {
		// arrange
		// ...

		// act
		input := []string{"hello", "world"}
		result := Some(input)

		// assert
		// - optional equality
		expected := Option[[]string]{value: new([]string)}; *expected.value = []string{"hello", "world"}
		require.Equal(t, expected, result)
		// - check if it is a Some
		require.True(t, result.IsSome())
		// - check unwrapped value
		expectedUnwrap := []string{"hello", "world"}
		require.Equal(t, expectedUnwrap, result.Unwrap())
	})

	t.Run("Some - []int", func(t *testing.T) {
		// arrange
		// ...

		// act
		input := []int{42, 42}
		result := Some(input)

		// assert
		// - optional equality
		expected := Option[[]int]{value: new([]int)}; *expected.value = []int{42, 42}
		require.Equal(t, expected, result)
		// - check if it is a Some
		require.True(t, result.IsSome())
		// - check unwrapped value
		expectedUnwrap := []int{42, 42}
		require.Equal(t, expectedUnwrap, result.Unwrap())
	})

	t.Run("Some - []float64", func(t *testing.T) {
		// arrange
		// ...

		// act
		input := []float64{42.0, 42.0}
		result := Some(input)

		// assert
		// - optional equality
		expected := Option[[]float64]{value: new([]float64)}; *expected.value = []float64{42.0, 42.0}
		require.Equal(t, expected, result)
		// - check if it is a Some
		require.True(t, result.IsSome())
		// - check unwrapped value
		expectedUnwrap := []float64{42.0, 42.0}
		require.Equal(t, expectedUnwrap, result.Unwrap())
	})

	t.Run("Some - []bool", func(t *testing.T) {
		// arrange
		// ...

		// act
		input := []bool{true, true}
		result := Some(input)

		// assert
		// - optional equality
		expected := Option[[]bool]{value: new([]bool)}; *expected.value = []bool{true, true}
		require.Equal(t, expected, result)
		// - check if it is a Some
		require.True(t, result.IsSome())
		// - check unwrapped value
		expectedUnwrap := []bool{true, true}
		require.Equal(t, expectedUnwrap, result.Unwrap())
	})

	t.Run("Some - struct", func(t *testing.T) {
		// arrange
		type testStruct struct {
			a int
			b string
		}
		
		// act
		input := testStruct{a: 42, b: "hello"}
		result := Some(input)

		// assert
		// - optional equality
		expected := Option[testStruct]{value: new(testStruct)}; *expected.value = testStruct{a: 42, b: "hello"}
		require.Equal(t, expected, result)
		// - check if it is a Some
		require.True(t, result.IsSome())
		// - check unwrapped value
		expectedUnwrap := testStruct{a: 42, b: "hello"}
		require.Equal(t, expectedUnwrap, result.Unwrap())
	})
}

// TestOption_None tests the None constructor.
func TestOption_None(t *testing.T) {
	t.Run("None - string", func(t *testing.T) {
		// arrange
		// ...

		// act
		result := None[string]()

		// assert
		var v string
		defer func() {
			if r := recover(); r != nil {
				require.Equal(t, "", v)
				require.Equal(t, "unable to unwrap None", r)
			}
		}()
		v = result.Unwrap()
	})

	t.Run("None - int", func(t *testing.T) {
		// arrange
		// ...

		// act
		result := None[int]()

		// assert
		var v int
		defer func() {
			if r := recover(); r != nil {
				require.Equal(t, 0, v)
				require.Equal(t, "unable to unwrap None", r)
			}
		}()
		v = result.Unwrap()
	})

	t.Run("None - float64", func(t *testing.T) {
		// arrange
		// ...

		// act
		result := None[float64]()

		// assert
		var v float64
		defer func() {
			if r := recover(); r != nil {
				require.Equal(t, 0.0, v)
				require.Equal(t, "unable to unwrap None", r)
			}
		}()
		v = result.Unwrap()
	})

	t.Run("None - bool", func(t *testing.T) {
		// arrange
		// ...

		// act
		result := None[bool]()

		// assert
		var v bool
		defer func() {
			if r := recover(); r != nil {
				require.Equal(t, false, v)
				require.Equal(t, "unable to unwrap None", r)
			}
		}()
		v = result.Unwrap()
	})

	t.Run("None - []string", func(t *testing.T) {
		// arrange
		// ...

		// act
		result := None[[]string]()

		// assert
		var v []string
		defer func() {
			if r := recover(); r != nil {
				require.Equal(t, []string(nil), v)
				require.Equal(t, "unable to unwrap None", r)
			}
		}()
		v = result.Unwrap()
	})

	t.Run("None - []int", func(t *testing.T) {
		// arrange
		// ...

		// act
		result := None[[]int]()

		// assert
		var v []int
		defer func() {
			if r := recover(); r != nil {
				require.Equal(t, []int(nil), v)
				require.Equal(t, "unable to unwrap None", r)
			}
		}()
		v = result.Unwrap()
	})

	t.Run("None - []float64", func(t *testing.T) {
		// arrange
		// ...

		// act
		result := None[[]float64]()

		// assert
		var v []float64
		defer func() {
			if r := recover(); r != nil {
				require.Equal(t, []float64(nil), v)
				require.Equal(t, "unable to unwrap None", r)
			}
		}()
		v = result.Unwrap()
	})

	t.Run("None - []bool", func(t *testing.T) {
		// arrange
		// ...

		// act
		result := None[[]bool]()

		// assert
		var v []bool
		defer func() {
			if r := recover(); r != nil {
				require.Equal(t, []bool(nil), v)
				require.Equal(t, "unable to unwrap None", r)
			}
		}()
		v = result.Unwrap()
	})

	t.Run("None - struct", func(t *testing.T) {
		// arrange
		type testStruct struct {}

		// act
		result := None[testStruct]()

		// assert
		var v testStruct
		defer func() {
			if r := recover(); r != nil {
				require.Equal(t, testStruct{}, v)
				require.Equal(t, "unable to unwrap None", r)
			}
		}()
		v = result.Unwrap()
	})
}