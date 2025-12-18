package str

// String is an immutable wrapper for fluent string operations.
type String struct {
	s string
}

// Of wraps a raw string with fluent helpers.
func Of(s string) String {
	return String{s: s}
}

// String returns the underlying raw string value.
func (s String) String() string {
	return s.s
}

// GoString allows %#v formatting to print the raw string.
func (s String) GoString() string {
	return s.s
}
