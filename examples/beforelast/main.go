// Command beforelast is generated as a standalone program so the documented BeforeLast example can be run directly.
package main

import "github.com/goforj/str/v2"

// main keeps this generated example directly runnable with go run.
func main() {
	// BeforeLast returns the substring before the last occurrence of sep.
	// If sep is empty or not found, the original string is returned.
	// Similar: Before and AfterLast.

	// Example: slice before last separator
	v := str.Of("pkg/path/file.txt").BeforeLast("/").String()
	println(v)
	// #string pkg/path
}
