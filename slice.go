package errors

import "strings"

// Slice is a slice of errors
type Slice []error

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
func ToSlice(err error, additionalErrors ...error) Slice {
	slice, ok := err.(Slice)
	if !ok {
		return Slice(append([]error{err}, additionalErrors...))
	}
	if len(additionalErrors) == 0 {
		return slice
	}
	return append(slice, additionalErrors...)
}

// MergeSlice merges one or more errors to a slice
func MergeSlice(err error, additionalErrors ...error) Slice {
	slice := ToSlice(err)
	for _, e := range additionalErrors {
		slice = append(slice, ToSlice(e)...)
	}
	return slice
}
