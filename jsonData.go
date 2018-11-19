/*
Package parsify is a declarative, extensible JSON parser

© 2018-present Harald Rudell <harald.rudell@gmail.com> (http://www.haraldrudell.com)
This source code is licensed under the ISC-style license found in the LICENSE file in the root directory of this source tree.
*/
package parsify

import (
	"encoding/json"
	"errors"
	"io"
)

// ParseByteStream data from a reader as JSON, the reader is typically resp.body from a http request
func ParseByteStream(reader io.Reader) (*JSONData, error) {
	data := JSONData{}
	e := json.NewDecoder(reader).Decode(&data.value)
	return &data, e
}

// ParseBytes coverts a utf-8 slice of bytes to objects
func ParseBytes(bytes []byte) (*JSONData, error) {
	data := JSONData{}
	return &data, json.Unmarshal(bytes, &data.value)
}

// JSONData hold parsed JSON data
type JSONData struct {
	value *interface{} // points to parsed JSO data, perhaps large
}

// Type determines the JSON type of current value
func (j *JSONData) Type() JSONType {
	if j == nil || j.value == nil { // no data
		return JSONnil
	}
	value := *j.value
	if value == nil {
		return JSONnull
	}
	if _, ok := value.(bool); ok {
		return JSONboolean
	}
	if _, ok := value.(float64); ok {
		return JSONnumber
	}
	if _, ok := value.(string); ok {
		return JSONstring
	}
	if _, ok := value.([]interface{}); ok {
		return JSONarray
	}
	if _, ok := value.(map[string]interface{}); ok {
		return JSONobject
	}
	panic(errors.New("JSONData corrupt"))
}

// IsNull examines the current value
func (j *JSONData) IsNull() bool {
	return j != nil && *j.value == nil
}

// Boolean examines the current value
func (j *JSONData) Boolean() *bool {
	if j != nil && j.value != nil {
		if bo, ok := (*j.value).(bool); ok {
			return &bo
		}
	}
	return nil
}

// Number examines the current value
func (j *JSONData) Number() *float64 {
	if j != nil && j.value != nil {
		if float, ok := (*j.value).(float64); ok {
			return &float
		}
	}
	return nil
}

func (j *JSONData) String() *string {
	if j != nil && j.value != nil {
		if string, ok := (*j.value).(string); ok {
			return &string
		}
	}
	return nil
}

func (j *JSONData) getArray() []interface{} {
	if j != nil && j.value != nil {
		if arr, ok := (*j.value).([]interface{}); ok {
			return arr
		}
	}
	return nil
}

// ArrayLength examines the current value
func (j *JSONData) ArrayLength() *int {
	if arr := j.getArray(); arr != nil {
		n := len(arr)
		return &n
	}
	return nil
}

// ArrayAt examines the current value
func (j *JSONData) ArrayAt(i int) *JSONData {
	if arr := j.getArray(); arr != nil {
		if i >= 0 && i < len(arr) {
			el := arr[i]
			return &JSONData{&el}
		}
	}
	return nil
}

// EnterArray examines the current value
func (j *JSONData) EnterArray(i int) *JSONData {
	if arr := j.getArray(); arr != nil {
		if i >= 0 && i < len(arr) {
			el := arr[i]
			j.value = &el
			return j
		}
	}
	return nil
}

func (j *JSONData) getObject() map[string]interface{} {
	if j != nil && j.value != nil {
		if mapp, ok := (*j.value).(map[string]interface{}); ok {
			return mapp
		}
	}
	return nil
}

// ObjectKeys examines the current value
func (j *JSONData) ObjectKeys() []string {
	if mapp := j.getObject(); mapp != nil {
		keys := make([]string, 0, len(mapp))
		i := 0
		for key := range mapp {
			keys[i] = key
			i++
		}
		return keys
	}
	return nil
}

// EnterObject zooms in to an object property value of the JSON
func (j *JSONData) EnterObject(key string) *JSONData {
	if mapp := j.getObject(); mapp != nil {
		if v, ok := mapp[key]; ok {
			j.value = &v
			return j
		}
	}
	return nil
}

// ObjectValue examines the current value
func (j *JSONData) ObjectValue(key string) *JSONData {
	if mapp := j.getObject(); mapp != nil {
		if v, ok := mapp[key]; ok {
			return &JSONData{&v}
		}
	}
	return nil
}
