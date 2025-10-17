package nullable

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// UInt64 is a container for uint64 type that provides nullable semantics without using pointers.
type UInt64 struct {
	IsPresent bool
	Value     uint64
}

// IsDefined determines whether this field should be included in the json output, if it has the omitempty tag.
func (v UInt64) IsDefined() bool {
	return v.IsPresent
}

// Get returns the value if it is not null, otherwise it returns the given default value.
func (v UInt64) Get(value uint64) uint64 {
	if v.IsPresent {
		return v.Value
	} else {
		return value
	}
}

// Set stores the value and sets it as not null.
func (v *UInt64) Set(value uint64) {
	v.IsPresent = true
	v.Value = value
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v UInt64) MarshalEasyJSON(w *jwriter.Writer) {
	if v.IsPresent {
		w.Uint64(v.Value)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *UInt64) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = UInt64{}
	} else {
		v.Value = l.Uint64()
		v.IsPresent = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v UInt64) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *UInt64) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}
