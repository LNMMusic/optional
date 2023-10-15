package optional

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

// UnmarshalBSONValue indicates how to unmarshal a bson value into an Option.
func (o *Option[T]) UnmarshalBSONValue(t bsontype.Type, data []byte) (err error) {
	// encapsulate data into bson.RawValue
	rv := bson.RawValue{Type: t, Value: data}

	// unmarshal bson.RawValue into o.value
	err = rv.Unmarshal(&o.value)
	return
}

// MarshalBSONValue indicates how to marshal an Option into a bson value.
func (o Option[T]) MarshalBSONValue() (t bsontype.Type, data []byte, err error) {
	// marshal o.value into bson.RawValue
	t, data, err = bson.MarshalValue(o.value)
	return
}