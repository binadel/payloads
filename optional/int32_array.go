package optional

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Int32Array is an optional array of int32 for providing optional semantics without using pointers.
type Int32Array struct {
	isDefined bool
	Value     []int32
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v Int32Array) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the isDefined to true.
func (v *Int32Array) SetDefined() {
	v.isDefined = true
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Int32Array) MarshalEasyJSON(w *jwriter.Writer) {
	w.RawByte('[')
	for i, item := range v.Value {
		if i > 0 {
			w.RawByte(',')
		}
		w.Int32(item)
	}
	w.RawByte(']')
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Int32Array) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Int32Array{}
	} else {
		v.Value = make([]int32, 0)
		l.Delim('[')
		for !l.IsDelim(']') {
			var item int32
			if l.IsNull() {
				l.Skip()
			} else {
				item = l.Int32()
			}
			v.Value = append(v.Value, item)
			l.WantComma()
		}
		l.Delim(']')
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Int32Array) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Int32Array) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
