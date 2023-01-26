package main

import (
	"runtime"

	"github.com/songx23/golang-present/slices/examples/memory-leak/log"
)

func main() {
	var fbs [][]byte
	for i := 0; i < 1_000; i++ {
		b := make([]byte, 1_000_000)
		fiveBts := getFirstFiveBytes(b)
		fbs = append(fbs, fiveBts)
	}
	runtime.GC()
	log.PrintAlloc()
	runtime.KeepAlive(fbs)
}

func getFirstFiveBytes(bts []byte) []byte {
	return bts[:5]
}
