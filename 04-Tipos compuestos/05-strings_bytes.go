package main

import "fmt"

func stringsBytes() {
	// var my_string string = "Hola Mundo"
	// var my_byte byte = my_string[0]
	// fmt.Println(my_byte)

	// // var s2 string = my_string[3:7]
	// var s3 string = my_string[:7]
	// fmt.Println(s3)

	// Strings a slices
	var s string = "Hola Mundo"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}
