package optional

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// UIntArray is a container for uint slice type that provides optional semantics without using pointers.
type UIntArray struct {
	isDefined bool
	Value     []uint
}

// IsDefined determines whether this field should be included in the json output, if it has the omitempty tag.
func (v UIntArray) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the field to defined, see IsDefined.
func (v *UIntArray) SetDefined() {
	v.isDefined = true
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v UIntArray) MarshalEasyJSON(w *jwriter.Writer) {
	if v.Value == nil {
		w.RawString("null")
	} else {
		w.RawByte('[')
		for i, item := range v.Value {
			if i > 0 {
				w.RawByte(',')
			}
			w.Uint(item)
		}
		w.RawByte(']')
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *UIntArray) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = UIntArray{}
	} else {
		v.Value = make([]uint, 0)
		l.Delim('[')
		for !l.IsDelim(']') {
			var item uint
			if l.IsNull() {
				l.Skip()
			} else {
				item = l.Uint()
			}
			v.Value = append(v.Value, item)
			l.WantComma()
		}
		l.Delim(']')
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v UIntArray) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *UIntArray) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
