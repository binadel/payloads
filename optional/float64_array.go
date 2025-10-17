package optional

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Float64Array is an optional array of float64 that provides optional semantics.
type Float64Array struct {
	isDefined bool
	Value     []float64
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v Float64Array) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the isDefined to true.
func (v *Float64Array) SetDefined() {
	v.isDefined = true
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Float64Array) MarshalEasyJSON(w *jwriter.Writer) {
	if v.Value == nil {
		w.RawString("null")
	} else {
		w.RawByte('[')
		for i, item := range v.Value {
			if i > 0 {
				w.RawByte(',')
			}
			w.Float64(item)
		}
		w.RawByte(']')
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Float64Array) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Float64Array{}
	} else {
		v.Value = make([]float64, 0)
		l.Delim('[')
		for !l.IsDelim(']') {
			var item float64
			if l.IsNull() {
				l.Skip()
			} else {
				item = l.Float64()
			}
			v.Value = append(v.Value, item)
			l.WantComma()
		}
		l.Delim(']')
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Float64Array) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Float64Array) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
