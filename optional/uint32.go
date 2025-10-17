package optional

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// UInt32 is an optional and nullable uint32 type that provides optional semantics without using pointers.
type UInt32 struct {
	isDefined bool
	IsPresent bool
	Value     uint32
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v UInt32) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the isDefined to true.
func (v *UInt32) SetDefined() {
	v.isDefined = true
}

// Get returns the value if not null; otherwise returns the default given.
func (v UInt32) Get(value uint32) uint32 {
	if v.IsPresent {
		return v.Value
	} else {
		return value
	}
}

// Set stores the value and sets it as not null.
func (v *UInt32) Set(value uint32) {
	v.IsPresent = true
	v.Value = value
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v UInt32) MarshalEasyJSON(w *jwriter.Writer) {
	if v.IsPresent {
		w.Uint32(v.Value)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *UInt32) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = UInt32{}
	} else {
		v.Value = l.Uint32()
		v.IsPresent = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v UInt32) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *UInt32) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
