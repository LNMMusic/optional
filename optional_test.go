package optional

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, output, result)
		assert.True(t, result.IsSome())
		v, err := result.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, input, v)
	})

	t.Run("Some - int", func(t *testing.T) {
		// arrange
		var input int = 1
		var output = Option[int]{value: &input}

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
		var output = Option[float64]{value: &input}

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
		var output = Option[bool]{value: &input}

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
		var output = Option[[]string]{value: &input}

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
		var output = Option[[]int]{value: &input}

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
		var output = Option[[]float64]{value: &input}

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
		var output = Option[[]bool]{value: &input}

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
		var output = Option[testStruct]{value: &input}

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

// TestOption_None tests the None constructor.
func TestOption_None(t *testing.T) {
	t.Run("None - string", func(t *testing.T) {
		// arrange
		var output = Option[string]{value: nil}

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
		var output = Option[int]{value: nil}

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
		var output = Option[float64]{value: nil}

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
		var output = Option[bool]{value: nil}

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
		var output = Option[[]string]{value: nil}

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
		var output = Option[[]int]{value: nil}

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
		var output = Option[[]float64]{value: nil}

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
		var output = Option[[]bool]{value: nil}

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
		var output = Option[testStruct]{value: nil}

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

// TestOption_IsSome tests the IsSome method.
func TestOption_Unmarshal(t *testing.T) {
	// string
	t.Run("JSON - string", func(t *testing.T) {
		// arrange
		data := []byte(`"hello"`)
		option := Option[string]{value: nil}

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

	// int
	t.Run("JSON - int", func(t *testing.T) {
		// arrange
		data := []byte(`1`)
		option := Option[int]{value: nil}

		// act
		err := json.Unmarshal(data, &option)

		// assert
		assert.NoError(t, err)
		assert.True(t, option.IsSome())
		v, err := option.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, 1, v)
	})
	t.Run("JSON - int null", func(t *testing.T) {
		// arrange
		data := []byte(`null`)
		option := Some(1)

		// act
		err := json.Unmarshal(data, &option)

		// assert
		assert.NoError(t, err)
		assert.False(t, option.IsSome())
		v, err := option.Unwrap()
		assert.Equal(t, 0, v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
	})

	// float64
	t.Run("JSON - float64", func(t *testing.T) {
		// arrange
		data := []byte(`1.0`)
		option := Option[float64]{value: nil}

		// act
		err := json.Unmarshal(data, &option)

		// assert
		assert.NoError(t, err)
		assert.True(t, option.IsSome())
		v, err := option.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, 1.0, v)
	})
	t.Run("JSON - float64 null", func(t *testing.T) {
		// arrange
		data := []byte(`null`)
		option := Some(1.0)

		// act
		err := json.Unmarshal(data, &option)

		// assert
		assert.NoError(t, err)
		assert.False(t, option.IsSome())
		v, err := option.Unwrap()
		assert.Equal(t, 0.0, v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
	})

	// bool
	t.Run("JSON - bool", func(t *testing.T) {
		// arrange
		data := []byte(`true`)
		option := Option[bool]{value: nil}

		// act
		err := json.Unmarshal(data, &option)

		// assert
		assert.NoError(t, err)
		assert.True(t, option.IsSome())
		v, err := option.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, true, v)
	})
	t.Run("JSON - bool null", func(t *testing.T) {
		// arrange
		data := []byte(`null`)
		option := Some(true)

		// act
		err := json.Unmarshal(data, &option)

		// assert
		assert.NoError(t, err)
		assert.False(t, option.IsSome())
		v, err := option.Unwrap()
		assert.Equal(t, false, v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
	})

	// []string
	t.Run("JSON - []string", func(t *testing.T) {
		// arrange
		data := []byte(`["hello","world"]`)
		option := Option[[]string]{value: nil}

		// act
		err := json.Unmarshal(data, &option)

		// assert
		assert.NoError(t, err)
		assert.True(t, option.IsSome())
		v, err := option.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, []string{"hello", "world"}, v)
	})
	t.Run("JSON - []string null", func(t *testing.T) {
		// arrange
		data := []byte(`null`)
		option := Some([]string{"hello", "world"})

		// act
		err := json.Unmarshal(data, &option)

		// assert
		assert.NoError(t, err)
		assert.False(t, option.IsSome())
		v, err := option.Unwrap()
		assert.Equal(t, []string(nil), v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
	})

	// []int
	t.Run("JSON - []int", func(t *testing.T) {
		// arrange
		data := []byte(`[1,2,3]`)
		option := Option[[]int]{value: nil}

		// act
		err := json.Unmarshal(data, &option)

		// assert
		assert.NoError(t, err)
		assert.True(t, option.IsSome())
		v, err := option.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, v)
	})
	t.Run("JSON - []int null", func(t *testing.T) {
		// arrange
		data := []byte(`null`)
		option := Some([]int{1, 2, 3})

		// act
		err := json.Unmarshal(data, &option)

		// assert
		assert.NoError(t, err)
		assert.False(t, option.IsSome())
		v, err := option.Unwrap()
		assert.Equal(t, []int(nil), v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
	})

	// []float64
	t.Run("JSON - []float64", func(t *testing.T) {
		// arrange
		data := []byte(`[1.0,2.0,3.0]`)
		option := Option[[]float64]{value: nil}

		// act
		err := json.Unmarshal(data, &option)

		// assert
		assert.NoError(t, err)
		assert.True(t, option.IsSome())
		v, err := option.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, []float64{1.0, 2.0, 3.0}, v)
	})
	t.Run("JSON - []float64 null", func(t *testing.T) {
		// arrange
		data := []byte(`null`)
		option := Some([]float64{1.0, 2.0, 3.0})

		// act
		err := json.Unmarshal(data, &option)

		// assert
		assert.NoError(t, err)
		assert.False(t, option.IsSome())
		v, err := option.Unwrap()
		assert.Equal(t, []float64(nil), v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
	})

	// []bool
	t.Run("JSON - []bool", func(t *testing.T) {
		// arrange
		data := []byte(`[true,false]`)
		option := Option[[]bool]{value: nil}

		// act
		err := json.Unmarshal(data, &option)

		// assert
		assert.NoError(t, err)
		assert.True(t, option.IsSome())
		v, err := option.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, []bool{true, false}, v)
	})
	t.Run("JSON - []bool null", func(t *testing.T) {
		// arrange
		data := []byte(`null`)
		option := Some([]bool{true, false})

		// act
		err := json.Unmarshal(data, &option)

		// assert
		assert.NoError(t, err)
		assert.False(t, option.IsSome())
		v, err := option.Unwrap()
		assert.Equal(t, []bool(nil), v)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnwrapNone)
	})

	// structered data (all types)
	t.Run("JSON - structered data - value to null", func(t *testing.T) {
		// arrange
		type testStruct struct {
			FirstName 	Option[string]
			Age 		Option[int]
			Height 		Option[float64]
			Licensed 	Option[bool]
			Pets 		Option[[]string]
			Numbers 	Option[[]int]
			Location 	Option[[]float64]
			Bools		Option[[]bool]			
		}
		ts := testStruct{
			FirstName: Some("Mary"),
			Age: Some(20),
			Height: Some(1.7),
			Licensed: Some(true),
			Pets: Some([]string{"dog", "cat"}),
			Numbers: Some([]int{1, 2, 3}),
			Location: Some([]float64{1.0, 2.0, 3.0}),
			Bools: Some([]bool{true, false}),
		}

		data := []byte(`
			{
				"FirstName": null,
				"Age": null,
				"Height": null,
				"Licensed": null,
				"Pets": null,
				"Numbers": null,
				"Location": null,
				"Bools": null
			}
		`)

		// act
		err := json.Unmarshal(data, &ts)

		// assert
		assert.NoError(t, err)
		assert.False(t, ts.FirstName.IsSome())
		assert.False(t, ts.Age.IsSome())
		assert.False(t, ts.Height.IsSome())
		assert.False(t, ts.Licensed.IsSome())
		assert.False(t, ts.Pets.IsSome())
		assert.False(t, ts.Numbers.IsSome())
		assert.False(t, ts.Location.IsSome())
		assert.False(t, ts.Bools.IsSome())
	})

	t.Run("JSON - structered data - null to value", func(t *testing.T) {
		// arrange
		type testStruct struct {
			FirstName 	Option[string]
			Age 		Option[int]
			Height 		Option[float64]
			Licensed 	Option[bool]
			Pets 		Option[[]string]
			Numbers 	Option[[]int]
			Location 	Option[[]float64]
			Bools		Option[[]bool]
		}
		ts := testStruct{
			FirstName: None[string](),
			Age: None[int](),
			Height: None[float64](),
			Licensed: None[bool](),
			Pets: None[[]string](),
			Numbers: None[[]int](),
			Location: None[[]float64](),
			Bools: None[[]bool](),
		}

		data := []byte(`
			{
				"FirstName": "Mary",
				"Age": 20,
				"Height": 1.7,
				"Licensed": true,
				"Pets": ["dog", "cat"],
				"Numbers": [1, 2, 3],
				"Location": [1.0, 2.0, 3.0],
				"Bools": [true, false]
			}
		`)

		// act
		err := json.Unmarshal(data, &ts)

		// assert
		assert.NoError(t, err)
		assert.True(t, ts.FirstName.IsSome()); v1, err1 := ts.FirstName.Unwrap(); assert.NoError(t, err1); assert.Equal(t, "Mary", v1)
		assert.True(t, ts.Age.IsSome()); v2, err2 := ts.Age.Unwrap(); assert.NoError(t, err2); assert.Equal(t, 20, v2)
		assert.True(t, ts.Height.IsSome()); v3, err3 := ts.Height.Unwrap(); assert.NoError(t, err3); assert.Equal(t, 1.7, v3)
		assert.True(t, ts.Licensed.IsSome()); v4, err4 := ts.Licensed.Unwrap(); assert.NoError(t, err4); assert.Equal(t, true, v4)
		assert.True(t, ts.Pets.IsSome()); v5, err5 := ts.Pets.Unwrap(); assert.NoError(t, err5); assert.Equal(t, []string{"dog", "cat"}, v5)
		assert.True(t, ts.Numbers.IsSome()); v6, err6 := ts.Numbers.Unwrap(); assert.NoError(t, err6); assert.Equal(t, []int{1, 2, 3}, v6)
		assert.True(t, ts.Location.IsSome()); v7, err7 := ts.Location.Unwrap(); assert.NoError(t, err7); assert.Equal(t, []float64{1.0, 2.0, 3.0}, v7)
		assert.True(t, ts.Bools.IsSome()); v8, err8 := ts.Bools.Unwrap(); assert.NoError(t, err8); assert.Equal(t, []bool{true, false}, v8)
	})
}

// TestOption_IsNone tests the IsNone method.
func TestOption_Marshal(t *testing.T) {
	type input struct {value any}
	type output struct {data []byte; err error}
	type testCase struct {
		title  string
		input  input
		output output
	}

	cases := []testCase{
		// string
		{
			title: "JSON - string",
			input: input{value: Some("hello")},
			output: output{data: []byte(`"hello"`), err: nil},
		},
		{
			title: "JSON - string null",
			input: input{value: None[string]()},
			output: output{data: []byte(`null`), err: nil},
		},

		// int
		{
			title: "JSON - int",
			input: input{value: Some(1)},
			output: output{data: []byte(`1`), err: nil},
		},
		{
			title: "JSON - int null",
			input: input{value: None[int]()},
			output: output{data: []byte(`null`), err: nil},
		},

		// float64
		{
			title: "JSON - float64",
			input: input{value: Some(1.0)},
			output: output{data: []byte(`1`), err: nil},
		},
		{
			title: "JSON - float64 null",
			input: input{value: None[float64]()},
			output: output{data: []byte(`null`), err: nil},
		},

		// bool
		{
			title: "JSON - bool",
			input: input{value: Some(true)},
			output: output{data: []byte(`true`), err: nil},
		},
		{
			title: "JSON - bool null",
			input: input{value: None[bool]()},
			output: output{data: []byte(`null`), err: nil},
		},

		// []string
		{
			title: "JSON - []string",
			input: input{value: Some([]string{"hello", "world"})},
			output: output{data: []byte(`["hello","world"]`), err: nil},
		},
		{
			title: "JSON - []string null",
			input: input{value: None[[]string]()},
			output: output{data: []byte(`null`), err: nil},
		},

		// []int
		{
			title: "JSON - []int",
			input: input{value: Some([]int{1, 2, 3})},
			output: output{data: []byte(`[1,2,3]`), err: nil},
		},
		{
			title: "JSON - []int null",
			input: input{value: None[[]int]()},
			output: output{data: []byte(`null`), err: nil},
		},

		// []float64
		{
			title: "JSON - []float64",
			input: input{value: Some([]float64{1.0, 2.0, 3.0})},
			output: output{data: []byte(`[1,2,3]`), err: nil},
		},
		{
			title: "JSON - []float64 null",
			input: input{value: None[[]float64]()},
			output: output{data: []byte(`null`), err: nil},
		},

		// []bool
		{
			title: "JSON - []bool",
			input: input{value: Some([]bool{true, false})},
			output: output{data: []byte(`[true,false]`), err: nil},
		},
		{
			title: "JSON - []bool null",
			input: input{value: None[[]bool]()},
			output: output{data: []byte(`null`), err: nil},
		},

		// structered data (all types)
		{
			title: "JSON - structered data",
			input: input{value: struct {
				FirstName 	Option[string]
				Age 		Option[int]
				Height 		Option[float64]
				Licensed 	Option[bool]
				Pets 		Option[[]string]
				Numbers 	Option[[]int]
				Location 	Option[[]float64]
				Bools		Option[[]bool]
			}{
				FirstName: Some("Mary"),
				Age: Some(20),
				Height: Some(1.7),
				Licensed: Some(true),
				Pets: Some([]string{"dog", "cat"}),
				Numbers: Some([]int{1, 2, 3}),
				Location: Some([]float64{1.0, 2.0, 3.0}),
				Bools: Some([]bool{true, false}),
			}},
			output: output{data: []byte(`{"FirstName":"Mary","Age":20,"Height":1.7,"Licensed":true,"Pets":["dog","cat"],"Numbers":[1,2,3],"Location":[1,2,3],"Bools":[true,false]}`), err: nil},
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			// arrange
			// ...

			// act
			data, err := json.Marshal(c.input.value)

			// assert
			assert.Equal(t, c.output.data, data)
			assert.ErrorIs(t, err, c.output.err)
		})
	}
}