package nullable

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Int32 is a nullable int32 type that provides optional semantics without using pointers.
type Int32 struct {
	IsPresent bool
	Value     int32
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v Int32) IsDefined() bool {
	return v.IsPresent
}

// Get returns the value if not null; otherwise returns the default given.
func (v Int32) Get(value int32) int32 {
	if v.IsPresent {
		return v.Value
	} else {
		return value
	}
}

// Set stores the value and sets it as not null.
func (v *Int32) Set(value int32) {
	v.IsPresent = true
	v.Value = value
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Int32) MarshalEasyJSON(w *jwriter.Writer) {
	if v.IsPresent {
		w.Int32(v.Value)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Int32) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Int32{}
	} else {
		v.IsPresent = true
		v.Value = l.Int32()
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Int32) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Int32) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
