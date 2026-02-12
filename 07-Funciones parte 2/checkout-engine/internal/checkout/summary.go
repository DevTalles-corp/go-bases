package checkout

import (
	"fmt"
	"time"
)

func PrintHeader(title string) {
	fmt.Printf("\n=== %s ===\n", title)
}

func PrintDivider() {
	fmt.Println("--------------------------")
}

func PrintKV(key string, value any) {
	fmt.Printf("%-12s : %v\n", key, value)
}
func PrintKV2(key string, value any, extra any) {
	fmt.Printf("%-12s : %v %v\n", key, value, extra)
}

func Track(name string) func() {
	start := time.Now()
	return func() {
		PrintKV("Time:"+name, time.Since(start))
	}
}
