package errors

import (
	"encoding/json"
	"strings"
)

// Slice is a slice of errors
type Slice []*Error

func (s Slice) Error() string {
	builder := &strings.Builder{}
	for _, err := range s {
		builder.WriteString(err.Error())
		builder.WriteRune('\n')
	}

	return strings.TrimRight(builder.String(), "\n")
}

// NewSlice creates a new empty errors slice
func NewSlice() Slice {
	return make(Slice, 0)
}

// IsSlice checks if an error is a errors.Slice
func IsSlice(err error) bool {
	_, ok := err.(Slice)
	return ok
}

// Add adds an error to the slice
func (s Slice) Add(errs ...error) Slice {
	return MergeSlice(s, errs...)
}

// ToSlice converts an error to a errors.Slice
// this function reminds me about the good old haskel days :D
func ToSlice(err error, additionalErrors ...error) Slice {
	if len(additionalErrors) == 0 {
		if s, ok := err.(Slice); ok {
			return s
		}
		return Slice{New(err)}
	}
	return append(Slice{New(err)}, ToSlice(additionalErrors[0], additionalErrors[1:]...)...)
}

// MergeSlice merges one or more errors to a slice
func MergeSlice(err error, additionalErrors ...error) Slice {
	slice := ToSlice(err)
	for _, e := range additionalErrors {
		slice = append(slice, ToSlice(e)...)
	}
	return slice
}

// MarshalJSON implements the json.Marshaler interface
func (s Slice) MarshalJSON() ([]byte, error) {
	return json.Marshal([]*Error(s))
}
