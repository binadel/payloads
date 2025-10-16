package optionull

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// BoolArray is an optional array of bool for providing optional semantics without using pointers.
type BoolArray struct {
	isDefined bool
	Value     []bool
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v BoolArray) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the isDefined to true.
func (v *BoolArray) SetDefined() {
	v.isDefined = true
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v BoolArray) MarshalEasyJSON(w *jwriter.Writer) {
	if v.Value == nil {
		w.RawString("null")
	} else {
		w.RawByte('[')
		for i, item := range v.Value {
			if i > 0 {
				w.RawByte(',')
			}
			w.Bool(item)
		}
		w.RawByte(']')
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *BoolArray) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = BoolArray{}
	} else {
		v.Value = make([]bool, 0)
		l.Delim('[')
		for !l.IsDelim(']') {
			var item bool
			if l.IsNull() {
				l.Skip()
			} else {
				item = l.Bool()
			}
			v.Value = append(v.Value, item)
			l.WantComma()
		}
		l.Delim(']')
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v BoolArray) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *BoolArray) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
