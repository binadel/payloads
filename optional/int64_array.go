package optional

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Int64Array is an optional array of int64 that provides optional semantics.
type Int64Array struct {
	isDefined bool
	Value     []int64
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v Int64Array) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the isDefined to true.
func (v *Int64Array) SetDefined() {
	v.isDefined = true
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Int64Array) MarshalEasyJSON(w *jwriter.Writer) {
	if v.Value == nil {
		w.RawString("null")
	} else {
		w.RawByte('[')
		for i, item := range v.Value {
			if i > 0 {
				w.RawByte(',')
			}
			w.Int64(item)
		}
		w.RawByte(']')
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Int64Array) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Int64Array{}
	} else {
		v.Value = make([]int64, 0)
		l.Delim('[')
		for !l.IsDelim(']') {
			var item int64
			if l.IsNull() {
				l.Skip()
			} else {
				item = l.Int64()
			}
			v.Value = append(v.Value, item)
			l.WantComma()
		}
		l.Delim(']')
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Int64Array) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Int64Array) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
