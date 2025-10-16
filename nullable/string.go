package nullable

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// String is a nullable string type for providing optional semantics without using pointers.
type String struct {
	IsPresent bool
	Value     string
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v String) IsDefined() bool {
	return v.IsPresent
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v String) MarshalEasyJSON(w *jwriter.Writer) {
	if v.IsPresent {
		w.String(v.Value)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *String) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = String{}
	} else {
		v.Value = l.String()
		v.IsPresent = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v String) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *String) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
