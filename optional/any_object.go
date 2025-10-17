package optional

import "encoding/json"

// AnyObject is a container for struct type that provides optional semantics without using pointers.
// It uses encoding/json for marshaling and unmarshaling.
type AnyObject[T any] struct {
	isDefined bool
	Value     T
}

// IsDefined determines whether this field should be included in the json output, if it has the omitempty tag.
func (v AnyObject[T]) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the field to defined, see IsDefined.
func (v *AnyObject[T]) SetDefined() {
	v.isDefined = true
}

// MarshalJSON implements a standard json marshaler interface.
func (v AnyObject[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Value)
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *AnyObject[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &v.Value)
}
