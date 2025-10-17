package optional

import (
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Object is a container for struct type that provides optional semantics without using pointers.
// The generic argument T must be of type pointer to any struct
// that implements easyjson marshaler and unmarshaler interfaces.
type Object[T easyjson.MarshalerUnmarshaler] struct {
	isDefined bool
	Value     T
	New       func() T
}

// IsDefined determines whether this field should be included in the json output, if it has the omitempty tag.
func (v Object[T]) IsDefined() bool {
	return v.isDefined
}

// SetDefined is the setter for isDefined, see IsDefined.
func (v *Object[T]) SetDefined(isDefined bool) {
	v.isDefined = isDefined
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
				panic("Cannot instantiate generic type from nil constructor, set New function to define the constructor")
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
