package main

import "fmt"

func main() {
	src := []int{0, 1, 2}
	var dst []int
	copy(dst, src)
	fmt.Println(dst)
}
