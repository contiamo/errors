package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToMap(t *testing.T) {
	errMap := ToMap(New("err1"), "field1")
	assert.Equal(t, errMap["field1"].Error(), "err1")
	errMap2 := ToMap(New("err1"))
	assert.Equal(t, errMap2[""].Error(), "err1")
}

func TestMapAdd(t *testing.T) {
	errMap := NewMap()
	errMap.Add("field1", New("err1"))
	errMap.Add("field2", New("err2"))
	assert.Equal(t, len(errMap), 2)
}

func TestIsMap(t *testing.T) {
	m := NewMap()
	e := New("err1")
	assert.True(t, IsMap(m))
	assert.False(t, IsMap(e))
}
