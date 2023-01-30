package main

import "fmt"

func main() {
	src := []int{0, 1, 2}
	dst := make([]int, 3)
	copy(dst, src)
	fmt.Println(dst)
}
