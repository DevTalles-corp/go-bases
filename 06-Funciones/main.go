package main

import "fmt"

func main() {
	ok := PrintOK("Todo esta en orden")
	PrintOK("Todo cool")
	PrintOK("Funcionó bien")

	fmt.Println(ok)

}

func PrintOK(msg string) string {
	fmt.Printf("OK - %s\n", msg)
	return msg
}
