package optional

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// UInt16Array is an optional array of uint16 for providing optional semantics without using pointers.
type UInt16Array struct {
	isDefined bool
	Value     []uint16
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v UInt16Array) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the isDefined to true.
func (v *UInt16Array) SetDefined() {
	v.isDefined = true
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v UInt16Array) MarshalEasyJSON(w *jwriter.Writer) {
	w.RawByte('[')
	for i, item := range v.Value {
		if i > 0 {
			w.RawByte(',')
		}
		w.Uint16(item)
	}
	w.RawByte(']')
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *UInt16Array) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = UInt16Array{}
	} else {
		v.Value = make([]uint16, 0)
		l.Delim('[')
		for !l.IsDelim(']') {
			var item uint16
			if l.IsNull() {
				l.Skip()
			} else {
				item = l.Uint16()
			}
			v.Value = append(v.Value, item)
			l.WantComma()
		}
		l.Delim(']')
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v UInt16Array) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *UInt16Array) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
