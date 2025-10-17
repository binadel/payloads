package optional

import (
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Array is an optional and nullable array type for providing optional semantics.
// The generic argument should be of type pointer to any struct
// that implements easyjson marshaler and unmarshaler interfaces.
type Array[T easyjson.MarshalerUnmarshaler] struct {
	isDefined bool
	Value     []T
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v Array[T]) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the isDefined to true.
func (v *Array[T]) SetDefined() {
	v.isDefined = true
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Array[T]) MarshalEasyJSON(w *jwriter.Writer) {
	if v.Value == nil {
		w.RawString("null")
	} else {
		w.RawByte('[')
		for i, item := range v.Value {
			if i > 0 {
				w.RawByte(',')
			}
			item.MarshalEasyJSON(w)
		}
		w.RawByte(']')
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Array[T]) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Array[T]{}
	} else {
		l.Delim('[')
		v.Value = make([]T, 0)
		for !l.IsDelim(']') {
			var item T
			if l.IsNull() {
				l.Skip()
			} else {
				item.UnmarshalEasyJSON(l)
			}
			v.Value = append(v.Value, item)
			l.WantComma()
		}
		l.Delim(']')
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Array[T]) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Array[T]) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
