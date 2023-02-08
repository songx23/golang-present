package main

import "fmt"

func main() {
	var s []string
	// empty: true | nil: true  | len: 0 | cap: 0

	s = []string(nil)
	// empty: true | nil: true  | len: 0 | cap: 0

	s = []string{}
	// empty: true | nil: false | len: 0 | cap: 0

	s = make([]string, 0)
	// empty: true | nil: false | len: 0 | cap: 0

	s = make([]string, 1, 10)
	// len: 1 | cap: 10

	fmt.Println(s) // ignore this line, it's here just to keep the compiler happy
}
