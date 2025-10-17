package nullable

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// UInt8 is a nullable uint8 type that provides optional semantics without using pointers.
type UInt8 struct {
	IsPresent bool
	Value     uint8
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v UInt8) IsDefined() bool {
	return v.IsPresent
}

// Get returns the value if not null; otherwise returns the default given.
func (v UInt8) Get(value uint8) uint8 {
	if v.IsPresent {
		return v.Value
	} else {
		return value
	}
}

// Set stores the value and sets it as not null.
func (v *UInt8) Set(value uint8) {
	v.IsPresent = true
	v.Value = value
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v UInt8) MarshalEasyJSON(w *jwriter.Writer) {
	if v.IsPresent {
		w.Uint8(v.Value)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *UInt8) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = UInt8{}
	} else {
		v.Value = l.Uint8()
		v.IsPresent = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v UInt8) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *UInt8) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
