package optional

import (
	"fmt"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Object is an optional and nullable object type for providing optional semantics.
// The generic argument should be of type pointer to any struct
// that implements easyjson marshaler and unmarshaler interfaces.
type Object[T easyjson.MarshalerUnmarshaler] struct {
	isDefined bool
	Value     T
	New       func() T
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v Object[T]) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the isDefined to true.
func (v *Object[T]) SetDefined() {
	v.isDefined = true
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Object[T]) MarshalEasyJSON(w *jwriter.Writer) {
	if any(v.Value) == nil {
		w.RawString("null")
	} else {
		v.Value.MarshalEasyJSON(w)
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Object[T]) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Object[T]{}
	} else {
		if any(v.Value) == nil {
			if v.New == nil {
				l.AddError(fmt.Errorf("optional.Object: Value is nil and New is not set"))
				l.Skip()
				return
			}
			v.Value = v.New()
		}
		v.Value.UnmarshalEasyJSON(l)
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Object[T]) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Object[T]) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
