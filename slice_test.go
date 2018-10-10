package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceMerging(t *testing.T) {
	err1 := New("err 1")
	err2 := New("err 2")
	assert.NotNil(t, err1)
	assert.NotNil(t, err2)
	combined := MergeSlice(err1, err2)
	assert.NotNil(t, combined)
	assert.Equal(t, len(ToSlice(combined)), 2)
	err3 := New("err 3")
	combined = MergeSlice(combined, err3)
	assert.Equal(t, len(ToSlice(combined)), 3)
	assert.True(t, IsSlice(combined))
}

func TestToSlice(t *testing.T) {
	err1 := New("err 1")
	errorSliceFromError := ToSlice(err1)
	errorSliceFromSlice := ToSlice(errorSliceFromError)
	assert.Equal(t, errorSliceFromError, errorSliceFromSlice)
	errorSliceFromMultipleErrors := ToSlice(New("err1"), New("err2"), New("err3"))
	assert.Equal(t, len(errorSliceFromMultipleErrors), 3)
}

func TestIsSlice(t *testing.T) {
	err1 := New("err 1")
	errorSlice := ToSlice(err1)
	assert.True(t, IsSlice(errorSlice))
	assert.False(t, IsSlice(err1))
}

func TestSliceErrorFunction(t *testing.T) {
	err1 := New("err1")
	err2 := New("err2")
	err := ToSlice(err1, err2)
	errMsg := err.Error()
	assert.Equal(t, errMsg, "err1\nerr2")
}

func TestNewSlice(t *testing.T) {
	err := NewSlice()
	assert.Equal(t, len(err), 0)
	err = MergeSlice(err, New("foo"))
}
