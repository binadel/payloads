package optional

import (
	"encoding/json"
)

// ArrayAny is an optional and nullable array container for any type using encoding/json.
// Null elements are skipped and recorded as errors.
type ArrayAny[T any] struct {
	isDefined bool
	Value     []T
	New       func() []T
}

func (v ArrayAny[T]) IsDefined() bool {
	return v.isDefined
}
func (v *ArrayAny[T]) SetDefined() {
	v.isDefined = true
}

func (v ArrayAny[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Value)
}

func (v *ArrayAny[T]) UnmarshalJSON(data []byte) error {
	if v.New == nil {
		panic("Cannot instantiate Array[T] from nil constructor, use New to define the constructor")
	}
	v.Value = v.New()
	return json.Unmarshal(data, &v.Value)
}
