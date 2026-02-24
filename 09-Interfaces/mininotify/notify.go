package main

import (
	"context"
	"fmt"
)

type EmailSenderFake struct {
	Sent int
}

// Service
type Service struct {
	sender *EmailSenderFake // Acoplamiento
}

func (sender *EmailSenderFake) Send(ctx context.Context, to Email, body string) error {
	_ = ctx
	sender.Sent++
	fmt.Println("EMAIL SENT: ", sender.Sent, "TO: ", to)
	fmt.Println("BODY: ", body)
	return nil
}

func NewService(sender *EmailSenderFake) *Service {
	return &Service{sender: sender}
}

func (service *Service) NotifyPaymentDue(ctx context.Context, event Event) error {
	body := fmt.Sprintf("Tienes un pago pendiente de %s (event=%s)", event.Amount, event.ID)
	return service.sender.Send(ctx, event.Email, body)
}
