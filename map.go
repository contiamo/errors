package errors

import (
	"strings"
)

// Map is a map from string to error
type Map map[string]Slice

func (m Map) Error() string {
	builder := &strings.Builder{}
	for key, err := range m {
		builder.WriteString(key)
		builder.WriteString(": ")
		builder.WriteString(err.Error())
		builder.WriteRune('\n')
	}
	return strings.TrimRight(builder.String(), "\n")
}

// NewMap creates a new empty error map
func NewMap() Map {
	return make(Map)
}

// IsMap checks if an error is a error map
func IsMap(err error) bool {
	_, ok := err.(Map)
	return ok
}

// ToMap converts an error to a map error, key is optional and defaults to the empty string
func ToMap(err error, key ...string) Map {
	if m, ok := err.(Map); ok {
		return m
	}
	errorMap := make(Map)
	if len(key) > 0 {
		errorMap[key[0]] = ToSlice(err)
	} else {
		errorMap["error"] = ToSlice(err)
	}
	return errorMap
}

// Add sets an error in the map
func (m Map) Add(key string, err error) Map {
	m[key] = MergeSlice(m[key], err)
	return m
}

// MergeMap merges multiple error maps
func MergeMap(err1 Map, err2 Map) Map {
	res := NewMap()
	for k, v := range err1 {
		res[k] = v
	}
	for k, v := range err2 {
		res[k] = MergeSlice(res[k], v)
	}
	return res
}
