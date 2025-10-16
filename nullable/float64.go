package nullable

import (
	"fmt"

	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Float64 is a nullable float64 type for providing optional semantics without using pointers.
type Float64 struct {
	IsPresent bool
	Value     float64
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v Float64) IsDefined() bool {
	return v.IsPresent
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Float64) MarshalEasyJSON(w *jwriter.Writer) {
	if v.IsPresent {
		w.Float64(v.Value)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Float64) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Float64{}
	} else {
		v.Value = l.Float64()
		v.IsPresent = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Float64) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Float64) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}

// String implements a stringer interface using fmt.Sprint for the value.
func (v Float64) String() string {
	if !v.IsPresent {
		return "<null>"
	}
	return fmt.Sprint(v.Value)
}
