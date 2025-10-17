package optionull

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Bool is an optional and nullable bool type that provides optional semantics without using pointers.
type Bool struct {
	isDefined bool
	IsPresent bool
	Value     bool
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v Bool) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the isDefined to true.
func (v *Bool) SetDefined() {
	v.isDefined = true
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
