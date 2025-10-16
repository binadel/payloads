package nullable

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Int32 is a nullable int32 type for providing optional semantics without using pointers.
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
		v.Value = l.Int32()
		v.IsPresent = true
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
