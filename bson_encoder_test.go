package optional

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

// Tests for UnmarshalBSONValue
func TestUnmarshalBSONValue(t *testing.T) {
	t.Run("succeed to unmarshal - field none into string", func(t *testing.T) {
		// arrange
		type schema struct {
			Field Option[string] `bson:"field"`
		}

		// act
		s := schema{Field: None[string]()}
		bytes, err := bson.Marshal(bson.M{"field": ""})
		require.NoError(t, err)

		err = bson.Unmarshal(bytes, &s)

		// assert
		expectedSchema := schema{Field: Some[string]("")}
		require.NoError(t, err)
		require.Equal(t, expectedSchema, s)
	})

	t.Run("succeed to unmarshal - field none into none", func(t *testing.T) {
		// arrange
		type schema struct {
			Field Option[string] `bson:"field"`
		}

		// act
		s := schema{Field: None[string]()}
		bytes, err := bson.Marshal(bson.M{"field": nil})
		require.NoError(t, err)

		err = bson.Unmarshal(bytes, &s)

		// assert
		expectedSchema := schema{Field: None[string]()}
		require.NoError(t, err)
		require.Equal(t, expectedSchema, s)
	})

	t.Run("succeed to unmarshal - field some into string", func(t *testing.T) {
		// arrange
		type schema struct {
			Field Option[string] `bson:"field"`
		}

		// act
		s := schema{Field: Some[string]("")}
		bytes, err := bson.Marshal(bson.M{"field": "new"})
		require.NoError(t, err)

		err = bson.Unmarshal(bytes, &s)

		// assert
		expectedSchema := schema{Field: Some[string]("new")}
		require.NoError(t, err)
		require.Equal(t, expectedSchema, s)
	})

	t.Run("succeed to unmarshal - field some into none", func(t *testing.T) {
		// arrange
		type schema struct {
			Field Option[string] `bson:"field"`
		}

		// act
		s := schema{Field: Some[string]("")}
		bytes, err := bson.Marshal(bson.M{"field": nil})
		require.NoError(t, err)

		err = bson.Unmarshal(bytes, &s)

		// assert
		expectedSchema := schema{Field: None[string]()}
		require.NoError(t, err)
		require.Equal(t, expectedSchema, s)
	})
}

// Tests for MarshalBSONValue
func TestMarshalBSONValue(t *testing.T) {
	t.Run("succeed to marshal - field should get nil", func(t *testing.T) {
		// arrange
		type schema struct {
			Field Option[string] `bson:"field"`
		}

		// act
		s := schema{Field: None[string]()}
		bytes, err := bson.Marshal(s)


		// assert
		require.NoError(t, err)
		expectedBytes, err := bson.Marshal(bson.M{"field": nil})
		require.NoError(t, err)
		require.Equal(t, expectedBytes, bytes)
	})

	t.Run("succeed to marshal - field should get empty string", func(t *testing.T) {
		// arrange
		type schema struct {
			Field Option[string] `bson:"field"`
		}

		// act
		s := schema{Field: Some[string]("")}
		bytes, err := bson.Marshal(s)

		// assert
		require.NoError(t, err)
		expectedBytes, err := bson.Marshal(bson.M{"field": ""})
		require.NoError(t, err)
		require.Equal(t, expectedBytes, bytes)
	})

	t.Run("succeed to marshal - field should get string", func(t *testing.T) {
		// arrange
		type schema struct {
			Field Option[string] `bson:"field"`
		}

		// act
		s := schema{Field: Some[string]("new")}
		bytes, err := bson.Marshal(s)

		// assert
		require.NoError(t, err)
		expectedBytes, err := bson.Marshal(bson.M{"field": "new"})
		require.NoError(t, err)
		require.Equal(t, expectedBytes, bytes)
	})
}