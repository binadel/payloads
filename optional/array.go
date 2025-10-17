package optional

import (
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Array is a container for slice type that provides optional semantics without using pointers.
// The generic argument T must be of type pointer to any struct
// that implements easyjson marshaler and unmarshaler interfaces.
type Array[T easyjson.MarshalerUnmarshaler] struct {
	isDefined bool
	Value     []T
	New       func() T
}

// IsDefined determines whether this field should be included in the json output, if it has the omitempty tag.
func (v Array[T]) IsDefined() bool {
	return v.isDefined
}

// SetDefined is the setter for isDefined, see IsDefined.
func (v *Array[T]) SetDefined(isDefined bool) {
	v.isDefined = isDefined
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
			if any(item) == nil {
				w.RawString("null")
			} else {
				item.MarshalEasyJSON(w)
			}
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
		v.Value = make([]T, 0)
		l.Delim('[')
		for !l.IsDelim(']') {
			var item T
			if l.IsNull() {
				l.Skip()
			} else {
				if v.New == nil {
					panic("Cannot instantiate generic type from nil constructor, set New function to define the constructor")
				}
				item = v.New()
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
