package main

import "fmt"

func main() {
	email, _ := NewEmail("ricardo@correo.com")
	amount, _ := NewMoneyFromCents(45000)

	fmt.Println("Email: ", email)
	fmt.Println("Money: ", amount)
}
