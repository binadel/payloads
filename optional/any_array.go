package optional

import "encoding/json"

// AnyArray is a container for slice type that provides optional semantics without using pointers.
// It uses encoding/json for marshaling and unmarshaling.
type AnyArray[T any] struct {
	isDefined bool
	Value     []T
}

// IsDefined determines whether this field should be included in the json output, if it has the omitempty tag.
func (v AnyArray[T]) IsDefined() bool {
	return v.isDefined
}

// SetDefined is the setter for isDefined, see IsDefined.
func (v *AnyArray[T]) SetDefined(isDefined bool) {
	v.isDefined = isDefined
}

// MarshalJSON implements a standard json marshaler interface.
func (v AnyArray[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Value)
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *AnyArray[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &v.Value)
}
