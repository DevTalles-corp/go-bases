package main

import (
	"context"
	"fmt"
)

func main() {
	email, _ := NewEmail("ricardo@correo.com")
	amount, _ := NewMoneyFromCents(45000)

	fmt.Println("Email: ", email)
	fmt.Println("Money: ", amount)

	ev := NewPaymentDueEvent("evt_001", email, amount)
	fmt.Println(ev.Type.String(), ev.ID, ev.Amount)

	// Probando Send
	send := EmailSenderFake{}
	body := "Pago pendiente: " + ev.Amount.String()

	_ = send.Send(context.Background(), ev.Email, body)
	_ = send.Send(context.Background(), ev.Email, body)
	_ = send.Send(context.Background(), ev.Email, body)
	_ = send.Send(context.Background(), ev.Email, body)
}
