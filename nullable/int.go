package nullable

import (
	"fmt"

	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Int is a nullable int type for providing optional semantics without using pointers.
type Int struct {
	IsPresent bool
	Value     int
}

// IsDefined returns whether the value is defined.
// It is used by easyjson when the field has omitempty tag,
// to decide whether to include the field or not.
func (v Int) IsDefined() bool {
	return v.IsPresent
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Int) MarshalEasyJSON(w *jwriter.Writer) {
	if v.IsPresent {
		w.Int(v.Value)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Int) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Int{}
	} else {
		v.Value = l.Int()
		v.IsPresent = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Int) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Int) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}

// String implements a stringer interface using fmt.Sprint for the value.
func (v Int) String() string {
	if !v.IsPresent {
		return "<null>"
	}
	return fmt.Sprint(v.Value)
}
