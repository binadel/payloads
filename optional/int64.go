package optional

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Int64 is an optional int64 type that provides optional semantics without using pointers.
type Int64 struct {
	isDefined bool
	Value     int64
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v Int64) IsDefined() bool {
	return v.isDefined
}

// SetDefined sets the isDefined to true.
func (v *Int64) SetDefined() {
	v.isDefined = true
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Int64) MarshalEasyJSON(w *jwriter.Writer) {
	w.Int64(v.Value)
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Int64) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Int64{}
	} else {
		v.Value = l.Int64()
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Int64) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Int64) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
