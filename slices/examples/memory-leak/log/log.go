package log

import (
	"fmt"
	"runtime"
)

func PrintAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%.6f KB\n", float64(m.Alloc)/1_024)
}
