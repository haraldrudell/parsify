/*
Package parsify is a declarative, extensible JSON parser

© 2018-present Harald Rudell <harald.rudell@gmail.com> (http://www.haraldrudell.com)
This source code is licensed under the ISC-style license found in the LICENSE file in the root directory of this source tree.
*/
package parsify

// JSONType describes underlyig JSON data type
type JSONType uint

// constants
const (
	JSONnil     JSONType = 0 // no data
	JSONnull                 // nil
	JSONboolean              // bool
	JSONnumber               // float64
	JSONstring               // string
	JSONarray                // array: []interface{}
	JSONobject               // map[string]interface{}
)

func (j JSONType) String() string {
	return [...]string{"JSONnull", "JSONboolean", "JSONnumber", "JSONstring", "JSONarray", "JSONobject"}[j]
}
