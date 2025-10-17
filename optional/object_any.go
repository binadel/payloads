package optional

import (
	"encoding/json"
)

// ObjectAny is an optional and nullable object container for any type using encoding/json.
type ObjectAny[T any] struct {
	isDefined bool
	Value     T
	New       func() T
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v ObjectAny[T]) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the isDefined to true.
func (v *ObjectAny[T]) SetDefined() {
	v.isDefined = true
}

// MarshalJSON implements a standard json marshaler interface.
func (v ObjectAny[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Value)
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *ObjectAny[T]) UnmarshalJSON(data []byte) error {
	if v.New == nil {
		panic("Cannot instantiate Object[T] from nil constructor, use New to define the constructor")
	}
	v.Value = v.New()
	return json.Unmarshal(data, v.Value)
}
