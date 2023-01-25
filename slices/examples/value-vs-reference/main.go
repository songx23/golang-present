package main

import "fmt"

func main() {
	// ARRAY
	a1 := [3]int{1, 2, 3}
	var a2 [3]int
	a2 = a1
	a2[1] = 10
	fmt.Println(a1) // [1, 2, 3]
	fmt.Println(a2) // [1, 10, 3]
	// SLICE
	s1 := []int{1, 2, 3}
	var s2 []int
	s2 = s1
	a1[1] = 10
	fmt.Println(s1) // [1, 10, 3]
	fmt.Println(s2) // [1, 10, 3]
}
