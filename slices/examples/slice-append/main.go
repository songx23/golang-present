package main

import (
	"fmt"
	"runtime"
)

func main() {
	s1 := []int{1, 2, 3}
	// s1: [1, 2, 3]
	// len: 3 | cap: 3

	s2 := s1[1:2]
	// s2: [2]
	// len: 1 | cap: 2

	s3 := append(s2, 4)
	// s3: [2, 4]
	// len: 2 | cap: 2
	// what's s1?

	fmt.Println(s1)
	runtime.KeepAlive(s3) // Ignore this line, it's just there for compilation.
}
