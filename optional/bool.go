package optional

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Bool is an optional bool type for providing optional semantics without using pointers.
type Bool struct {
	isDefined bool
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
	w.Bool(v.Value)
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Bool) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Bool{}
	} else {
		v.Value = l.Bool()
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
