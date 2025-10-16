package optional

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// UInt8Array is an optional array of uint8 for providing optional semantics without using pointers.
type UInt8Array struct {
	isDefined bool
	Value     []uint8
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v UInt8Array) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the isDefined to true.
func (v *UInt8Array) SetDefined() {
	v.isDefined = true
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v UInt8Array) MarshalEasyJSON(w *jwriter.Writer) {
	w.RawByte('[')
	for i, item := range v.Value {
		if i > 0 {
			w.RawByte(',')
		}
		w.Uint8(item)
	}
	w.RawByte(']')
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *UInt8Array) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = UInt8Array{}
	} else {
		v.Value = make([]uint8, 0)
		l.Delim('[')
		for !l.IsDelim(']') {
			var item uint8
			if l.IsNull() {
				l.Skip()
			} else {
				item = l.Uint8()
			}
			v.Value = append(v.Value, item)
			l.WantComma()
		}
		l.Delim(']')
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v UInt8Array) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *UInt8Array) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
