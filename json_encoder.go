package optional

import "encoding/json"

// UnmarshalJSON indicates how to unmarshal a json value into an Option.
// unmarshalling (always work with reference)
// json.Unmarshal: always sends a ptr of the field
// - interface: always accepts ptr | no matter method receiver
//
// cases
// - method [ptr-receiver] -> works
// - method [non-ptr]      -> does not work (can't access to the field of &o.Value)
func (o *Option[T]) UnmarshalJSON(data []byte) (err error) {
	err = json.Unmarshal(data, &o.value)
	return
}

// MarshalJSON indicates how to marshal an Option into a json value.
// marshalling (always work with value)
// json.Marshal: always sends a non-ptr of the field
// - interface: not always accepts non-ptr
//              > case non-ptr: works
//              > case ptr: does not work (go rules)
//
// cases
// - method [non-ptr] -> works ()
func (o Option[T]) MarshalJSON() (data []byte, err error) {
	data, err = json.Marshal(o.value)
	return
}