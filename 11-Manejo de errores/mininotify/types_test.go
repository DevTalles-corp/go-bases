package main

import (
	"errors"
	"testing"
)

func TestNewEmailOK(test *testing.T) {
	email, err := NewEmail("RicArdo@Correo.com")
	if err != nil {
		test.Fatalf("Esperabamos un nil, obtuvimos: %v", err)
	}

	if email != "ricardo@correo.com" {
		test.Fatalf("No se esperaba el correo: %q", email)
	}
}

func TestNewEmailInvalid(test *testing.T) {
	_, err := NewEmail("not-an-email")
	if err == nil {
		test.Fatalf("Esperabamos un err, obtuvimos nil")
	}

	if !errors.Is(err, ErrInvalidEmail) {
		test.Fatalf("Esperabamos errors.Is(err, ErrInvalidEmail) = true, obtuvimos: %v", err)
	}

}

func TestNewMoneyFromCentsOK(test *testing.T) {
	amount, err := NewMoneyFromCents(45000)
	if err != nil {
		test.Fatalf("Esperabamos un nil, obtuvimos: %v", err)
	}

	if amount.String() != "$450.00" {
		test.Fatalf("No se esperaba el formato de monto: %s", amount.String())
	}
}

func TestNewMoneyFromCentsNegative(test *testing.T) {
	_, err := NewMoneyFromCents(-1)
	if err == nil {
		test.Fatalf("Experabamos un error y obtuvimos un nil")
	}

	if !errors.Is(err, ErrNegativeMoney) {
		test.Fatalf("Esperabamos un error errors.Is(err, ErrNegativeMoney) = true, obtuvimos: %v", err)
	}
}
