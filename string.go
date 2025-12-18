package str

// String is an immutable wrapper for fluent string operations.
// @group Fluent
type String struct {
	s string
}

// Of wraps a raw string with fluent helpers.
// @group Fluent
//
// Example: wrap raw string
//
//	v := str.Of("gopher")
//	godump.Dump(v.String())
//	// #string gopher
func Of(s string) String {
	return String{s: s}
}

// String returns the underlying raw string value.
// @group Fluent
//
// Example: unwrap to plain string
//
//	v := str.Of("go").String()
//	godump.Dump(v)
//	// #string go
func (s String) String() string {
	return s.s
}

// GoString allows %#v formatting to print the raw string.
// @group Fluent
//
// Example: fmt %#v uses GoString
//
//	v := str.Of("go")
//	godump.Dump(fmt.Sprintf("%#v", v))
//	// #string go
func (s String) GoString() string {
	return s.s
}
