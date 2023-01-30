package main

import "fmt"

func main() {
	var names = []string{"s1", "s2", "s3"}
	var slices = make([]*[]int, 0, 3)
	s1 := []int{1, 100, 10_000}
	slices = append(slices, &s1)
	log(names, slices)
	s2 := s1[1:2]
	slices = append(slices, &s2)
	log(names, slices)
	s3 := append(s2, 1_000_000)
	slices = append(slices, &s3)
	log(names, slices)
	s3 = append(s3, 100_000_000)
	log(names, slices)
}
func log(names []string, slices []*[]int) {
	for i, s := range slices {
		fmt.Printf("slice: %s\nvalue:%v\nlen: %d, cap: %d\n", names[i], *s, len(*s), cap(*s))
	}
	fmt.Println("---------")
}
