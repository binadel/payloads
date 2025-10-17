package optional

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Int16 is a container for int16 type that provides optional semantics without using pointers.
type Int16 struct {
	isDefined bool
	IsPresent bool
	Value     int16
}

// IsDefined determines whether this field should be included in the json output, if it has the omitempty tag.
func (v Int16) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the field to defined, see IsDefined.
func (v *Int16) SetDefined() {
	v.isDefined = true
}

// Get returns the value if it is not null, otherwise it returns the given default value.
func (v Int16) Get(value int16) int16 {
	if v.IsPresent {
		return v.Value
	} else {
		return value
	}
}

// Set stores the value and sets it as not null.
func (v *Int16) Set(value int16) {
	v.IsPresent = true
	v.Value = value
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Int16) MarshalEasyJSON(w *jwriter.Writer) {
	if v.IsPresent {
		w.Int16(v.Value)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Int16) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Int16{}
	} else {
		v.Value = l.Int16()
		v.IsPresent = true
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
