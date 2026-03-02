package main

import (
	"errors"
	"fmt"
)

func main() {
	email, err := NewEmail("ricardo")
	if err != nil {
		if errors.Is(err, ErrInvalidEmail) {
			fmt.Println("Email inválido. Corrige el formato: ", err)
			return
		}
		fmt.Println("ERROR al crear el email: ", err)
		return
	}

	amount, err := NewMoneyFromCents(-45000)
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
