package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Printf("Length: %d, Capacity: %d\n", len(arr), cap(arr))

	s := make([]int, 3, 6)
	fmt.Printf("Length: %d, Capacity: %d\n", len(s), cap(s))
}
