package errors

import (
	"encoding/json"
	"io"

	yaml "gopkg.in/yaml.v2"
)

// MarshalToObject marshals to a map[string]interface{}, usable to marshal to different formats
func MarshalToObject(err error, additionalErrors ...error) map[string]interface{} {
	fieldErrors := NewMap()
	sliceErrors := NewSlice()
	errs := append([]error{err}, additionalErrors...)
	for _, err := range errs {
		if IsMap(err) {
			fieldErrors = MergeMap(fieldErrors, ToMap(err))
		} else {
			sliceErrors = sliceErrors.Add(err)
		}
	}
	obj := make(map[string]interface{})
	if len(fieldErrors) > 0 {
		m := make(map[string][]string)
		for k, v := range fieldErrors {
			for _, e := range v {
				m[k] = append(m[k], e.Error())
			}
		}
		obj["fieldErrors"] = m
	}
	if len(sliceErrors) > 0 {
		s := make([]string, len(sliceErrors))
		for idx, val := range sliceErrors {
			s[idx] = val.Error()
		}
		obj["errors"] = s
	}
	return obj
}

// MarshalToJSON an error to a json encoded error object
func MarshalToJSON(writer io.Writer, err error, additionalErrors ...error) error {
	obj := MarshalToObject(err, additionalErrors...)
	encoder := json.NewEncoder(writer)
	return encoder.Encode(obj)
}

// MarshalToYAML marshals an error to a yaml encoded error object
func MarshalToYAML(writer io.Writer, err error, additionalErrors ...error) error {
	obj := MarshalToObject(err, additionalErrors...)
	bs, err := yaml.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = writer.Write(bs)
	return err
}
