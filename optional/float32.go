package optional

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Float32 is an optional and nullable float32 type that provides optional semantics without using pointers.
type Float32 struct {
	isDefined bool
	IsPresent bool
	Value     float32
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v Float32) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the isDefined to true.
func (v *Float32) SetDefined() {
	v.isDefined = true
}

// Get returns the value if not null; otherwise returns the default given.
func (v Float32) Get(value float32) float32 {
	if v.IsPresent {
		return v.Value
	} else {
		return value
	}
}

// Set stores the value and sets it as not null.
func (v *Float32) Set(value float32) {
	v.IsPresent = true
	v.Value = value
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Float32) MarshalEasyJSON(w *jwriter.Writer) {
	if v.IsPresent {
		w.Float32(v.Value)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Float32) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Float32{}
	} else {
		v.Value = l.Float32()
		v.IsPresent = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Float32) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Float32) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
