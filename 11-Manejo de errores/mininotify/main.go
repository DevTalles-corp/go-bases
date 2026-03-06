package main

import (
	"errors"
	"fmt"
)

func main() {
	email, err := NewEmail("ricardo@example.com")
	if err != nil {
		if errors.Is(err, ErrInvalidEmail) {
			fmt.Println("Email inválido. Corrige el formato: ", err)
			return
		}
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

	if err := WriteReport("report.txt", "Reporte de la semana"); err != nil {
		fmt.Println("Al crear el reporte. Error: ", err)
	}

	if err := ShadowingError(); err != nil {
		fmt.Println("Shadowing error: ", err)
	} else {
		fmt.Println("ShadowingError retorno nil (TODO BIEN)")
	}

	if err := RunSafely(func() {
		_ = MustTemplate("payment_due")
	}); err != nil {
		fmt.Println("SAFE ERROR: ", err)
	}

}

func returnsTypedNil() Sender {
	var s *EmailSenderFake = nil
	return s
}
