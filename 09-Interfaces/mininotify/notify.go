package main

import (
	"context"
	"fmt"
)

type EmailSenderFake struct {
	Sent int
}

type Sender interface {
	Send(ctx context.Context, to Email, body string) error
}

// Service
// type Service struct {
// 	sender *EmailSenderFake // Acoplamiento
// }

// Nuevo service con interfaz
type Service struct {
	sender Sender
}

func (sender *EmailSenderFake) Send(ctx context.Context, to Email, body string) error {
	_ = ctx
	sender.Sent++
	fmt.Println("EMAIL SENT: ", sender.Sent, "TO: ", to)
	fmt.Println("BODY: ", body)
	return nil
}

func NewService(sender Sender) *Service {
	return &Service{sender: sender}
}

func (service *Service) NotifyPaymentDue(ctx context.Context, event Event) error {
	body := fmt.Sprintf("Tienes un pago pendiente de %s (event=%s)", event.Amount, event.ID)
	return service.sender.Send(ctx, event.Email, body)
}
