package str

import "encoding/base64"

// ToBase64 encodes the string using standard Base64.
// @group Encoding
//
// Example: base64 encode
//
//	v := str.Of("gopher").ToBase64().String()
//	godump.Dump(v)
//	// #string Z29waGVy
func (s String) ToBase64() String {
	return String{s: base64.StdEncoding.EncodeToString([]byte(s.s))}
}

// FromBase64 decodes a standard Base64 string.
// @group Encoding
//
// Example: base64 decode
//
//	v, err := str.Of("Z29waGVy").FromBase64()
//	godump.Dump(v.String(), err)
//	// #string gopher
//	// #error <nil>
func (s String) FromBase64() (String, error) {
	decoded, err := base64.StdEncoding.DecodeString(s.s)
	return String{s: string(decoded)}, err
}
