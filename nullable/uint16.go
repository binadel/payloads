package nullable

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// UInt16 is a nullable uint16 type that provides optional semantics without using pointers.
type UInt16 struct {
	IsPresent bool
	Value     uint16
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v UInt16) IsDefined() bool {
	return v.IsPresent
}

// Get returns the value if not null; otherwise returns the default given.
func (v UInt16) Get(value uint16) uint16 {
	if v.IsPresent {
		return v.Value
	} else {
		return value
	}
}

// Set stores the value and sets it as not null.
func (v *UInt16) Set(value uint16) {
	v.IsPresent = true
	v.Value = value
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v UInt16) MarshalEasyJSON(w *jwriter.Writer) {
	if v.IsPresent {
		w.Uint16(v.Value)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *UInt16) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = UInt16{}
	} else {
		v.Value = l.Uint16()
		v.IsPresent = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v UInt16) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *UInt16) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
