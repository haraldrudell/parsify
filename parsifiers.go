/*
Package parsify is a declarative, extensible JSON parser

© 2018-present Harald Rudell <harald.rudell@gmail.com> (http://www.haraldrudell.com)
This source code is licensed under the ISC-style license found in the LICENSE file in the root directory of this source tree.
*/
package parsify

import "fmt"

// Fn signature for JSON-parsing step functions
type Fn func(*Step, *Steps) error

// Fns is a map for providing groups of step functions
type Fns map[string]Fn

// DefaultFns some basic JSON parsing functions
var DefaultFns = Fns{
	"VerifyStringProperty": verifyStringProperty,
	"VerifyNumberProperty": verifyNumberProperty,
	"EnterKey":             enterKey,
	"StoreNumber":          storeNumber,
}

// ParsifyJSON run a Parsify list
func verifyStringProperty(s *Step, ps *Steps) error {
	pStr := ps.data.GetStringObjectValue(s.parameter)
	strValue, ok := s.value.(string)
	if pStr == nil || !ok || *pStr != strValue {
		return ps.getError(fmt.Sprintf("bad value for %s", s.parameter))
	}
	return nil
}

func verifyNumberProperty(s *Step, ps *Steps) error {
	pf := ps.data.GetNumberObjectValue(s.parameter)
	value, ok := s.value.(float64)
	if pf == nil || !ok || *pf != value {
		return ps.getError(fmt.Sprintf("bad value for %s", s.parameter))
	}
	return nil
}

func enterKey(s *Step, ps *Steps) error {
	pf := ps.data.EnterObject(s.parameter)
	if pf == nil {
		return ps.getError(fmt.Sprintf("key not found: '%s'", s.parameter))
	}
	return nil
}

func storeNumber(s *Step, ps *Steps) error {
	pf := ps.data.GetNumericStringObjectValue(s.parameter) // hex string to *uint64
	if pf != nil {
		ptr, ok := s.storeLocation.(*uint64)
		if ok {
			*ptr = *pf
			return nil
		}
	}
	return ps.getError(fmt.Sprintf("failed to store number: '%s'", s.parameter))
}
