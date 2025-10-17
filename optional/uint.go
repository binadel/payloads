package optional

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// UInt is an optional and nullable uint type that provides optional semantics without using pointers.
type UInt struct {
	isDefined bool
	IsPresent bool
	Value     uint
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v UInt) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the isDefined to true.
func (v *UInt) SetDefined() {
	v.isDefined = true
}

// Get returns the value if not null; otherwise returns the default given.
func (v UInt) Get(value uint) uint {
	if v.IsPresent {
		return v.Value
	} else {
		return value
	}
}

// Set stores the value and sets it as not null.
func (v *UInt) Set(value uint) {
	v.IsPresent = true
	v.Value = value
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v UInt) MarshalEasyJSON(w *jwriter.Writer) {
	if v.IsPresent {
		w.Uint(v.Value)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *UInt) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = UInt{}
	} else {
		v.Value = l.Uint()
		v.IsPresent = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v UInt) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *UInt) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
