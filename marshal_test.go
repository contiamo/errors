package errors

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalToObject(t *testing.T) {
	err1 := New("err1")
	err2 := New("err2")
	err3 := New("err3")
	err4 := New("err4")
	slice := ToSlice(err1, err2)
	m := ToMap(err3, "field1")
	obj := MarshalToObject(slice, m, err4)
	assert.Equal(t, 2, len(obj))
	assert.Equal(t, 3, len(obj["errors"].([]string)))
	assert.Equal(t, 1, len(obj["fieldErrors"].(map[string][]string)))
}

func TestMarshalToYAML(t *testing.T) {
	err1 := New("err1")
	err2 := New("err2")
	err3 := New("err3")
	err4 := New("err4")
	slice := ToSlice(err1, err2)
	m := ToMap(err3, "field1")
	buf := &bytes.Buffer{}
	MarshalToYAML(buf, slice, m, err4)
	s := buf.String()
	expected := `errors:
- err1
- err2
- err4
fieldErrors:
  field1:
  - err3
`
	assert.Equal(t, expected, s)
}

func TestMarshalToJSON(t *testing.T) {
	err1 := New("err1")
	err2 := New("err2")
	err3 := New("err3")
	err4 := New("err4")
	slice := ToSlice(err1, err2)
	m := ToMap(err3, "field1")
	buf := &bytes.Buffer{}
	MarshalToJSON(buf, slice, m, err4)
	s := buf.String()
	expected := `{"errors":["err1","err2","err4"],"fieldErrors":{"field1":["err3"]}}` + "\n"
	assert.Equal(t, expected, s)
}

func TestJSONMarshalerInterface(t *testing.T) {
	sliceMarshaler := json.Marshaler(ToSlice(New("err1")))
	bs, err := json.Marshal(sliceMarshaler)
	assert.NoError(t, err)
	assert.Equal(t, []byte(`["err1"]`), bs)
	mapMarshaler := json.Marshaler(Map{"f1": Slice{New("err1")}})
	bs, err = json.Marshal(mapMarshaler)
	assert.NoError(t, err)
	assert.Equal(t, []byte(`{"f1":["err1"]}`), bs)
}

func TestJSONUnmarshalerInterface(t *testing.T) {
	s := Slice{}
	bs := []byte(`["err1"]`)
	err := json.Unmarshal(bs, &s)
	assert.NoError(t, err)
	assert.Equal(t, "err1", s[0].Error())
}
