package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToMap(t *testing.T) {
	errMap := ToMap(New("err1"), "field1")
	assert.Equal(t, errMap["field1"].Error(), "err1")
	errMap2 := ToMap(New("err1"))
	assert.Equal(t, errMap2["error"].Error(), "err1")
}

func TestMapAdd(t *testing.T) {
	errMap := NewMap()
	errMap.Add("field1", New("err1"))
	errMap.Add("field2", New("err2"))
	errMap.Add("field2", New("err3"))
	assert.Equal(t, len(errMap), 2)
	assert.Equal(t, len(errMap["field2"]), 2)
}

func TestIsMap(t *testing.T) {
	m := NewMap()
	e := New("err1")
	assert.True(t, IsMap(m))
	assert.False(t, IsMap(e))
}

func TestWeirdMergeMap(t *testing.T) {
	m1 := Map{"f1": Slice{New("e1")}, "f2": Slice{New("e2")}}
	m2 := Map{"f1": Slice{New("e4")}, "f3": Slice{New("e3")}}
	m := MergeMap(m1, m2)
	assert.Equal(t, 2, len(m["f1"]))
	assert.Equal(t, 1, len(m["f2"]))
	assert.Equal(t, 1, len(m["f3"]))

}
