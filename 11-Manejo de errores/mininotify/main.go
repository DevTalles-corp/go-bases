package main

import (
	"fmt"
)

func main() {
	email, err := NewEmail("ricardo@correo.com")
	if err != nil {
		fmt.Println("ERROR al crear el email: ", err)
		return
	}

	amount, err := NewMoneyFromCents(45000)
	if err != nil {
		fmt.Println("ERROR al crear el monto: ", err)
		return
	}

	fmt.Println("Email: ", email)
	fmt.Println("Money: ", amount)

}

func returnsTypedNil() Sender {
	var s *EmailSenderFake = nil
	return s
}
