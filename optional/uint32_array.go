package optional

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// UInt32Array is an optional array of uint32 that provides optional semantics.
type UInt32Array struct {
	isDefined bool
	Value     []uint32
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v UInt32Array) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the isDefined to true.
func (v *UInt32Array) SetDefined() {
	v.isDefined = true
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v UInt32Array) MarshalEasyJSON(w *jwriter.Writer) {
	if v.Value == nil {
		w.RawString("null")
	} else {
		w.RawByte('[')
		for i, item := range v.Value {
			if i > 0 {
				w.RawByte(',')
			}
			w.Uint32(item)
		}
		w.RawByte(']')
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *UInt32Array) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = UInt32Array{}
	} else {
		v.Value = make([]uint32, 0)
		l.Delim('[')
		for !l.IsDelim(']') {
			var item uint32
			if l.IsNull() {
				l.Skip()
			} else {
				item = l.Uint32()
			}
			v.Value = append(v.Value, item)
			l.WantComma()
		}
		l.Delim(']')
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v UInt32Array) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *UInt32Array) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
