package optional

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

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
		require.NoError(t, err)
		require.True(t, option.IsSome())
		require.Equal(t, "hello", option.Unwrap())
	})
	t.Run("JSON - string null", func(t *testing.T) {
		// arrange
		data := []byte(`null`)
		option := Some("hello")

		// act
		err := json.Unmarshal(data, &option)

		// assert
		require.NoError(t, err)
		require.False(t, option.IsSome())
	})

	// int
	t.Run("JSON - int", func(t *testing.T) {
		// arrange
		data := []byte(`1`)
		option := Option[int]{value: nil}

		// act
		err := json.Unmarshal(data, &option)

		// assert
		require.NoError(t, err)
		require.True(t, option.IsSome())
		require.Equal(t, 1, option.Unwrap())
	})
	t.Run("JSON - int null", func(t *testing.T) {
		// arrange
		data := []byte(`null`)
		option := Some(1)

		// act
		err := json.Unmarshal(data, &option)

		// assert
		require.NoError(t, err)
		require.False(t, option.IsSome())
	})

	// float64
	t.Run("JSON - float64", func(t *testing.T) {
		// arrange
		data := []byte(`1.0`)
		option := Option[float64]{value: nil}

		// act
		err := json.Unmarshal(data, &option)

		// assert
		require.NoError(t, err)
		require.True(t, option.IsSome())
		require.Equal(t, 1.0, option.Unwrap())
	})
	t.Run("JSON - float64 null", func(t *testing.T) {
		// arrange
		data := []byte(`null`)
		option := Some(1.0)

		// act
		err := json.Unmarshal(data, &option)

		// assert
		require.NoError(t, err)
		require.False(t, option.IsSome())
	})

	// bool
	t.Run("JSON - bool", func(t *testing.T) {
		// arrange
		data := []byte(`true`)
		option := Option[bool]{value: nil}

		// act
		err := json.Unmarshal(data, &option)

		// assert
		require.NoError(t, err)
		require.True(t, option.IsSome())
		require.Equal(t, true, option.Unwrap())
	})
	t.Run("JSON - bool null", func(t *testing.T) {
		// arrange
		data := []byte(`null`)
		option := Some(true)

		// act
		err := json.Unmarshal(data, &option)

		// assert
		require.NoError(t, err)
		require.False(t, option.IsSome())
	})

	// []string
	t.Run("JSON - []string", func(t *testing.T) {
		// arrange
		data := []byte(`["hello","world"]`)
		option := Option[[]string]{value: nil}

		// act
		err := json.Unmarshal(data, &option)

		// assert
		require.NoError(t, err)
		require.True(t, option.IsSome())
		require.Equal(t, []string{"hello", "world"}, option.Unwrap())
	})
	t.Run("JSON - []string null", func(t *testing.T) {
		// arrange
		data := []byte(`null`)
		option := Some([]string{"hello", "world"})

		// act
		err := json.Unmarshal(data, &option)

		// assert
		require.NoError(t, err)
		require.False(t, option.IsSome())
	})

	// []int
	t.Run("JSON - []int", func(t *testing.T) {
		// arrange
		data := []byte(`[1,2,3]`)
		option := Option[[]int]{value: nil}

		// act
		err := json.Unmarshal(data, &option)

		// assert
		require.NoError(t, err)
		require.True(t, option.IsSome())
		require.Equal(t, []int{1, 2, 3}, option.Unwrap())
	})
	t.Run("JSON - []int null", func(t *testing.T) {
		// arrange
		data := []byte(`null`)
		option := Some([]int{1, 2, 3})

		// act
		err := json.Unmarshal(data, &option)

		// assert
		require.NoError(t, err)
		require.False(t, option.IsSome())
	})

	// []float64
	t.Run("JSON - []float64", func(t *testing.T) {
		// arrange
		data := []byte(`[1.0,2.0,3.0]`)
		option := Option[[]float64]{value: nil}

		// act
		err := json.Unmarshal(data, &option)

		// assert
		require.NoError(t, err)
		require.True(t, option.IsSome())
		require.Equal(t, []float64{1.0, 2.0, 3.0}, option.Unwrap())
	})
	t.Run("JSON - []float64 null", func(t *testing.T) {
		// arrange
		data := []byte(`null`)
		option := Some([]float64{1.0, 2.0, 3.0})

		// act
		err := json.Unmarshal(data, &option)

		// assert
		require.NoError(t, err)
		require.False(t, option.IsSome())
	})

	// []bool
	t.Run("JSON - []bool", func(t *testing.T) {
		// arrange
		data := []byte(`[true,false]`)
		option := Option[[]bool]{value: nil}

		// act
		err := json.Unmarshal(data, &option)

		// assert
		require.NoError(t, err)
		require.True(t, option.IsSome())
		require.Equal(t, []bool{true, false}, option.Unwrap())
	})
	t.Run("JSON - []bool null", func(t *testing.T) {
		// arrange
		data := []byte(`null`)
		option := Some([]bool{true, false})

		// act
		err := json.Unmarshal(data, &option)

		// assert
		require.NoError(t, err)
		require.False(t, option.IsSome())
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
		require.NoError(t, err)
		require.False(t, ts.FirstName.IsSome())
		require.False(t, ts.Age.IsSome())
		require.False(t, ts.Height.IsSome())
		require.False(t, ts.Licensed.IsSome())
		require.False(t, ts.Pets.IsSome())
		require.False(t, ts.Numbers.IsSome())
		require.False(t, ts.Location.IsSome())
		require.False(t, ts.Bools.IsSome())
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
		require.NoError(t, err)
		require.True(t, ts.FirstName.IsSome()); require.Equal(t, "Mary", ts.FirstName.Unwrap())
		require.True(t, ts.Age.IsSome()); require.Equal(t, 20, ts.Age.Unwrap())
		require.True(t, ts.Height.IsSome()); require.Equal(t, 1.7, ts.Height.Unwrap())
		require.True(t, ts.Licensed.IsSome()); require.Equal(t, true, ts.Licensed.Unwrap())
		require.True(t, ts.Pets.IsSome()); require.Equal(t, []string{"dog", "cat"}, ts.Pets.Unwrap())
		require.True(t, ts.Numbers.IsSome()); require.Equal(t, []int{1, 2, 3}, ts.Numbers.Unwrap())
		require.True(t, ts.Location.IsSome()); require.Equal(t, []float64{1.0, 2.0, 3.0}, ts.Location.Unwrap())
		require.True(t, ts.Bools.IsSome()); require.Equal(t, []bool{true, false}, ts.Bools.Unwrap())
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
			require.Equal(t, c.output.data, data)
			require.ErrorIs(t, err, c.output.err)
		})
	}
}