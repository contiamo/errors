package errors

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Error is a marshalable error
// note that you loose any extra information other than the actual message when marshaling/unmarshaling
type Error struct {
	err error
}

func (e *Error) Error() string {
	return e.err.Error()
}

// New creates a new marshalable error
// it may take a string or another error
func New(e interface{}) *Error {
	switch typedError := e.(type) {
	case error:
		return &Error{typedError}
	case string:
		return &Error{errors.New(typedError)}
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.err.Error())
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (e *Error) UnmarshalJSON(bs []byte) error {
	msg := ""
	if err := json.Unmarshal(bs, &msg); err != nil {
		return err
	}
	e.err = errors.New(msg)
	return nil
}
