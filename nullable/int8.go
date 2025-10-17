package nullable

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Int8 is a container for int8 type that provides nullable semantics without using pointers.
type Int8 struct {
	IsPresent bool
	Value     int8
}

// IsDefined determines whether this field should be included in the json output, if it has the omitempty tag.
func (v Int8) IsDefined() bool {
	return v.IsPresent
}

// Get returns the value if it is not null, otherwise it returns the given default value.
func (v Int8) Get(value int8) int8 {
	if v.IsPresent {
		return v.Value
	} else {
		return value
	}
}

// Set stores the value and sets it as not null.
func (v *Int8) Set(value int8) {
	v.IsPresent = true
	v.Value = value
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Int8) MarshalEasyJSON(w *jwriter.Writer) {
	if v.IsPresent {
		w.Int8(v.Value)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Int8) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Int8{}
	} else {
		v.Value = l.Int8()
		v.IsPresent = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Int8) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Int8) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
