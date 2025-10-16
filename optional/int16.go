package optional

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Int16 is an optional int16 type for providing optional semantics without using pointers.
type Int16 struct {
	isDefined bool
	Value     int16
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v Int16) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the isDefined to true.
func (v *Int16) SetDefined() {
	v.isDefined = true
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Int16) MarshalEasyJSON(w *jwriter.Writer) {
	w.Int16(v.Value)
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Int16) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Int16{}
	} else {
		v.Value = l.Int16()
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Int16) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Int16) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
