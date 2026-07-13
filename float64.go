package str

import "strconv"

// Float64 parses the string as a float64 using strconv.ParseFloat semantics.
// Similar: Bool and Int.
// @group Conversion
//
// Example: parse float64
//
//	v, err := str.Of("3.14").Float64()
//	fmt.Println(v, err == nil)
//	// #float64 3.14
//	// #bool true
func (s String) Float64() (float64, error) {
	return strconv.ParseFloat(s.s, 64)
}
