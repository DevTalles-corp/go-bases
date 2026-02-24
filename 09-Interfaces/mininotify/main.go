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
	// sender := &EmailSenderFake{}
	// svc := NewService(sender)
	// body := "Pago pendiente: " + ev.Amount.String()

	// _ = svc.NotifyPaymentDue(context.Background(), ev)
	// _ = send.Send(context.Background(), ev.Email, body)
	// _ = send.Send(context.Background(), ev.Email, body)
	// _ = send.Send(context.Background(), ev.Email, body)
	// _ = send.Send(context.Background(), ev.Email, body)

	base := &EmailSenderFake{}
	sender := LoggedSender{Sender: base, Log: ConsoleLogger{}}
	svc := NewService(sender)
	_ = svc.NotifyPaymentDue(context.Background(), ev)

	// Nil trap
	senderNil := returnsTypedNil()
	fmt.Println("senderNil == nil?", senderNil == nil)

}

func returnsTypedNil() Sender {
	var s *EmailSenderFake = nil
	return s
}
