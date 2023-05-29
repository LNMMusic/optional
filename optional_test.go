package optional

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOption_Some(t *testing.T) {
	t.Run("Some - string", func(t *testing.T) {
		// arrange
		var input string = "hello"
		var output = Option[string]{Value: &input}

		// act
		result := Some(input)

		// assert
		assert.Equal(t, output, result)
		assert.True(t, result.IsSome())
		v, err := result.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, input, v)
	})

	t.Run("Some - int", func(t *testing.T) {
		// arrange
		var input int = 1
		var output = Option[int]{Value: &input}

		// act
		result := Some(input)

		// assert
		assert.Equal(t, output, result)
		assert.True(t, result.IsSome())
		v, err := result.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, input, v)
	})

	t.Run("Some - float64", func(t *testing.T) {
		// arrange
		var input float64 = 1.0
		var output = Option[float64]{Value: &input}

		// act
		result := Some(input)

		// assert
		assert.Equal(t, output, result)
		assert.True(t, result.IsSome())
		v, err := result.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, input, v)
	})

	t.Run("Some - bool", func(t *testing.T) {
		// arrange
		var input bool = true
		var output = Option[bool]{Value: &input}

		// act
		result := Some(input)

		// assert
		assert.Equal(t, output, result)
		assert.True(t, result.IsSome())
		v, err := result.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, input, v)
	})

	t.Run("Some - []string", func(t *testing.T) {
		// arrange
		var input []string = []string{"hello", "world"}
		var output = Option[[]string]{Value: &input}

		// act
		result := Some(input)

		// assert
		assert.Equal(t, output, result)
		assert.True(t, result.IsSome())
		v, err := result.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, input, v)
	})

	t.Run("Some - []int", func(t *testing.T) {
		// arrange
		var input []int = []int{1, 2, 3}
		var output = Option[[]int]{Value: &input}

		// act
		result := Some(input)

		// assert
		assert.Equal(t, output, result)
		assert.True(t, result.IsSome())
		v, err := result.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, input, v)
	})

	t.Run("Some - []float64", func(t *testing.T) {
		// arrange
		var input []float64 = []float64{1.0, 2.0, 3.0}
		var output = Option[[]float64]{Value: &input}

		// act
		result := Some(input)

		// assert
		assert.Equal(t, output, result)
		assert.True(t, result.IsSome())
		v, err := result.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, input, v)
	})

	t.Run("Some - []bool", func(t *testing.T) {
		// arrange
		var input []bool = []bool{true, false}
		var output = Option[[]bool]{Value: &input}

		// act
		result := Some(input)

		// assert
		assert.Equal(t, output, result)
		assert.True(t, result.IsSome())
		v, err := result.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, input, v)
	})

	t.Run("Some - struct", func(t *testing.T) {
		// arrange
		type testStruct struct {
			value string
		}
		var input testStruct = testStruct{value: "hello"}
		var output = Option[testStruct]{Value: &input}

		// act
		result := Some(input)

		// assert
		assert.Equal(t, output, result)
		assert.True(t, result.IsSome())
		v, err := result.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, input, v)
	})
}

func TestOption_None(t *testing.T) {
	t.Run("None - string", func(t *testing.T) {
		// arrange
		var output = Option[string]{Value: nil}

		// act
		result := None[string]()

		// assert
		assert.Equal(t, output, result)
		assert.False(t, result.IsSome())
		v, err := result.Unwrap()
		assert.Equal(t, "", v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
	})

	t.Run("None - int", func(t *testing.T) {
		// arrange
		var output = Option[int]{Value: nil}

		// act
		result := None[int]()

		// assert
		assert.Equal(t, output, result)
		assert.False(t, result.IsSome())
		v, err := result.Unwrap()
		assert.Equal(t, 0, v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
	})

	t.Run("None - float64", func(t *testing.T) {
		// arrange
		var output = Option[float64]{Value: nil}

		// act
		result := None[float64]()

		// assert
		assert.Equal(t, output, result)
		assert.False(t, result.IsSome())
		v, err := result.Unwrap()
		assert.Equal(t, 0.0, v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
	})

	t.Run("None - bool", func(t *testing.T) {
		// arrange
		var output = Option[bool]{Value: nil}

		// act
		result := None[bool]()

		// assert
		assert.Equal(t, output, result)
		assert.False(t, result.IsSome())
		v, err := result.Unwrap()
		assert.Equal(t, false, v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
	})

	t.Run("None - []string", func(t *testing.T) {
		// arrange
		var output = Option[[]string]{Value: nil}

		// act
		result := None[[]string]()

		// assert
		assert.Equal(t, output, result)
		assert.False(t, result.IsSome())
		v, err := result.Unwrap()
		assert.Equal(t, []string(nil), v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
	})

	t.Run("None - []int", func(t *testing.T) {
		// arrange
		var output = Option[[]int]{Value: nil}

		// act
		result := None[[]int]()

		// assert
		assert.Equal(t, output, result)
		assert.False(t, result.IsSome())
		v, err := result.Unwrap()
		assert.Equal(t, []int(nil), v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
	})

	t.Run("None - []float64", func(t *testing.T) {
		// arrange
		var output = Option[[]float64]{Value: nil}

		// act
		result := None[[]float64]()

		// assert
		assert.Equal(t, output, result)
		assert.False(t, result.IsSome())
		v, err := result.Unwrap()
		assert.Equal(t, []float64(nil), v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
	})

	t.Run("None - []bool", func(t *testing.T) {
		// arrange
		var output = Option[[]bool]{Value: nil}

		// act
		result := None[[]bool]()

		// assert
		assert.Equal(t, output, result)
		assert.False(t, result.IsSome())
		v, err := result.Unwrap()
		assert.Equal(t, []bool(nil), v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
	})

	t.Run("None - struct", func(t *testing.T) {
		// arrange
		type testStruct struct {}
		var output = Option[testStruct]{Value: nil}

		// act
		result := None[testStruct]()

		// assert
		assert.Equal(t, output, result)
		assert.False(t, result.IsSome())
		v, err := result.Unwrap()
		assert.Equal(t, testStruct{}, v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
	})
}

func TestOption_Unmarshall(t *testing.T) {
	t.Run("JSON - string", func(t *testing.T) {
		// arrange
		data := []byte(`"hello"`)
		option := Option[string]{Value: nil}

		// act
		err := json.Unmarshal(data, &option)

		// assert
		assert.NoError(t, err)
		assert.True(t, option.IsSome())
		v, err := option.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, "hello", v)
	})

	t.Run("JSON - string null", func(t *testing.T) {
		// arrange
		data := []byte(`null`)
		option := Some("hello")

		// act
		err := json.Unmarshal(data, &option)

		// assert
		assert.NoError(t, err)
		assert.False(t, option.IsSome())
		v, err := option.Unwrap()
		assert.Equal(t, "", v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
	})

	t.Run("JSON - structered data", func(t *testing.T) {
		// arrange
		type testStruct struct {
			FirstName 	Option[string]
			LastName  	Option[string]
			Description Option[string]
		}
		ts := testStruct{
			FirstName: Some("Mary"),
			LastName:  Some("Jane"),
			Description: Some("This is a person"),
		}

		data := []byte(`{"FirstName":"John","LastName":null}`)

		// act
		err := json.Unmarshal(data, &ts)

		// assert
		assert.NoError(t, err)
		assert.True(t, ts.FirstName.IsSome())
		v, err := ts.FirstName.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, "John", v)
		assert.False(t, ts.LastName.IsSome())
		v, err = ts.LastName.Unwrap()
		assert.Equal(t, "", v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
		assert.True(t, ts.Description.IsSome())
		v, err = ts.Description.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, "This is a person", v)
	})
}

func TestOption_Marshall(t *testing.T) {
	t.Run("JSON - string", func(t *testing.T) {
		// arrange
		option := Some("hello")

		// act
		data, err := json.Marshal(option)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, []byte(`"hello"`), data)
	})

	t.Run("JSON - string null", func(t *testing.T) {
		// arrange
		option := None[string]()

		// act
		data, err := json.Marshal(option)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, []byte(`null`), data)
	})

	t.Run("JSON - structered data", func(t *testing.T) {
		// arrange
		type testStruct struct {
			FirstName 	Option[string]
			LastName  	Option[string]
			Description Option[string]
		}
		ts := testStruct{
			FirstName: Some("Mary"),
			LastName:  Some("Jane"),
			Description: None[string](),
		}

		// act
		data, err := json.Marshal(ts)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"FirstName":"Mary","LastName":"Jane","Description":null}`), data)
	})
}