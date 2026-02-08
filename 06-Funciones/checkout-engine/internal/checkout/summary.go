package checkout

import "fmt"

func PrintHeader(title string) {
	fmt.Printf("\n=== %s ===\n", title)
}

func PrintDivider() {
	fmt.Println("--------------------------")
}

func PrintKV(key string, value any) {
	fmt.Printf("%-12s : %v", key, value)
}
