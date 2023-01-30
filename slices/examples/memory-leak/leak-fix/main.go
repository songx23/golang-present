package main

import (
	"fmt"
	"runtime"
)

func main() {
	var fbs = make([][]byte, 1_000)
	for i := 0; i < 1_000; i++ {
		b := make([]byte, 1_000_000)
		fiveBts := getFirstFiveBytes(b)
		fbs[i] = fiveBts
	}
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%.2f MB\n", float64(m.Alloc)/1_024/1_000)
	runtime.KeepAlive(fbs)
}
func getFirstFiveBytes(bts []byte) []byte {
	res := make([]byte, 5)
	copy(res, bts)
	return res
}
