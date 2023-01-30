package main

import (
	"fmt"
	"runtime"
)

func main() {
	var fbs [][]byte
	for i := 0; i < 1_000; i++ {
		b := make([]byte, 1_000_000)
		fiveBts := getFirstFiveBytes(b)
		fbs = append(fbs, fiveBts)
	}
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%.6f KB\n", float64(m.Alloc)/1_024)
	runtime.KeepAlive(fbs)
}

func getFirstFiveBytes(bts []byte) []byte {
	return bts[:5]
}
