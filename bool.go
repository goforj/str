package str

import "strconv"

// Bool parses the string as a bool using strconv.ParseBool semantics.
// @group Conversion
//
// Example: parse bool
//
//	v, err := str.Of("true").Bool()
//	println(v, err == nil)
//	// #bool true
//	// #bool true
func (s String) Bool() (bool, error) {
	return strconv.ParseBool(s.s)
}
