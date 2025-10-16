package nullable

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Int64 is a nullable int64 type for providing optional semantics without using pointers.
type Int64 struct {
	IsPresent bool
	Value     int64
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v Int64) IsDefined() bool {
	return v.IsPresent
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Int64) MarshalEasyJSON(w *jwriter.Writer) {
	if v.IsPresent {
		w.Int64(v.Value)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Int64) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Int64{}
	} else {
		v.Value = l.Int64()
		v.IsPresent = true
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
