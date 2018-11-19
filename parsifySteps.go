/*
Package parsify is a declarative, extensible JSON parser

© 2018-present Harald Rudell <harald.rudell@gmail.com> (http://www.haraldrudell.com)
This source code is licensed under the ISC-style license found in the LICENSE file in the root directory of this source tree.
*/
package parsify

import (
	"fmt"
	"io"
)

// New creates a parsifier to be executed
func New(header string, steps []Step, fns Fns, data *JSONData) *Steps {
	allFns := make(Fns, len(DefaultFns)+len(fns))
	for _, mapp := range []Fns{DefaultFns, fns} {
		for key, value := range mapp {
			allFns[key] = value
		}
	}
	return &Steps{header, data, steps, 0, allFns}
}

// ParsifyReader instatiate and parsify
func (ps *Steps) ParsifyReader(reader io.Reader) (*JSONData, error) {
	data, e := ParseByteStream(reader)
	if e != nil {
		return nil, e
	}
	ps.data = data
	return data, ps.Parsify()
}

// ParsifyBytes parsifies from a utf-8 byte slice
func (ps *Steps) ParsifyBytes(bytes []byte) (*JSONData, error) {
	data, e := ParseBytes(bytes)
	if e != nil {
		return nil, e
	}
	ps.data = data
	return data, ps.Parsify()
}

// Steps instructions for parsing JSON
type Steps struct {
	heading string    // informative description of location for error messages
	data    *JSONData // the JSON data beig parsed
	steps   []Step    // steps to be executed
	no      int       // step number 1…
	fnMap   Fns       // map string to parisifer function
}

// Step instructions for parsing JSON
type Step struct {
	fn            string
	parameter     string
	value         interface{}
	storeLocation interface{} // pointer to something
}

// Parsify run a Parsify list
func (ps *Steps) Parsify() error {
	if ps.data == nil {
		return ps.getError("Parsify with data nil")
	}
	fnMap := ps.fnMap
	if fnMap == nil {
		fnMap = DefaultFns
	}
	for no, step := range ps.steps {
		ps.no = no + 1
		fn := fnMap[step.fn]
		if fn == nil {
			return ps.getError(fmt.Sprintf("unkown function: '%s'", step.fn))
		}
		if e := fn(&step, ps); e != nil {
			return e
		}
	}
	return nil
}

func (ps *Steps) getError(s string) error {
	return fmt.Errorf("%s step %d: %s", ps.heading, ps.no, s)
}
