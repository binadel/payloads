package nullable

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Bool is a container for bool type that provides nullable semantics without using pointers.
type Bool struct {
	IsPresent bool
	Value     bool
}

// IsDefined determines whether this field should be included in the json output, if it has the omitempty tag.
func (v Bool) IsDefined() bool {
	return v.IsPresent
}

// Get returns the value if it is not null, otherwise it returns the given default value.
func (v Bool) Get(value bool) bool {
	if v.IsPresent {
		return v.Value
	} else {
		return value
	}
}

// Set stores the value and sets it as not null.
func (v *Bool) Set(value bool) {
	v.IsPresent = true
	v.Value = value
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Bool) MarshalEasyJSON(w *jwriter.Writer) {
	if v.IsPresent {
		w.Bool(v.Value)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Bool) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Bool{}
	} else {
		v.Value = l.Bool()
		v.IsPresent = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Bool) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Bool) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
