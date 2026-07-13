package str

import "strconv"

// Int parses the string as a base-10 int using strconv.Atoi semantics.
// Similar: Bool and Float64.
// @group Conversion
//
// Example: parse int
//
//	v, err := str.Of("42").Int()
//	println(v, err == nil)
//	// #int 42
//	// #bool true
func (s String) Int() (int, error) {
	return strconv.Atoi(s.s)
}
