/*
Package parsify is a declarative, extensible JSON parser

© 2018-present Harald Rudell <harald.rudell@gmail.com> (http://www.haraldrudell.com)
This source code is licensed under the ISC-style license found in the LICENSE file in the root directory of this source tree.
*/
package parsify

import (
	"fmt"
	"strconv"
)

func (j *JSONData) getProperty(key string) *(interface{}) {
	if mapp := j.getObject(); mapp != nil {
		if v, ok := mapp[key]; ok {
			return &v
		}
	}
	return nil
}

// GetStringObjectValue ensure object and property to be string value
func (j *JSONData) GetStringObjectValue(key string) *string {
	pt := j.getProperty(key)
	if pt != nil {
		if stringValue, ok := (*pt).(string); ok {
			return &stringValue
		}
	}
	return nil
}

// GetNumberObjectValue ensure object and property to be string value
func (j *JSONData) GetNumberObjectValue(key string) *float64 {
	pt := j.getProperty(key)
	if pt != nil {
		if float, ok := (*pt).(float64); ok {
			return &float
		}
	}
	return nil
}

// GetNumericStringObjectValue gets timestamp, block number
func (j *JSONData) GetNumericStringObjectValue(key string) *uint64 {
	pt := j.getProperty(key)
	if pt != nil {
		if str, ok := (*pt).(string); ok {
			if num, e := strconv.ParseUint(str, 0, 64); e == nil {
				return &num
			}
		}
	}
	return nil
}

// Print what do I have?
func (j *JSONData) Print() {
	if j.value != nil {
		fmt.Printf("json: %T %#[1]v\n", *j.value)
	} else {
		fmt.Println("json: nil")
	}
}
