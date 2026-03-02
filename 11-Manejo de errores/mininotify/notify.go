package main

import (
	"context"
	"fmt"
)

type EmailSenderFake struct {
	Sent int
}

type Sender interface {
	Channel() string
	Send(ctx context.Context, to Email, body string) error
}

type Logger interface {
	Info(msg string)
}

// Service
// type Service struct {
// 	sender *EmailSenderFake // Acoplamiento
// }

// Nuevo service con interfaz
type Service struct {
	sender Sender
}

type ConsoleLogger struct{}

// Embedding
type LoggedSender struct {
	Sender
	Log Logger
}

func (sender *EmailSenderFake) Send(ctx context.Context, to Email, body string) error {
	_ = ctx
	sender.Sent++
	fmt.Println("EMAIL SENT: ", sender.Sent, "TO: ", to)
	fmt.Println("BODY: ", body)
	return nil
}

func (sender *EmailSenderFake) Channel() string {
	return "email_fake"
}

func NewService(sender Sender) *Service {
	return &Service{sender: sender}
}

func (service *Service) NotifyPaymentDue(ctx context.Context, event Event) error {
	body := fmt.Sprintf("[%s] Tienes un pago pendiente de %s (event=%s)", service.sender.Channel(), event.Amount, event.ID)
	return service.sender.Send(ctx, event.Email, body)
}

// Métodos para Logged
func (ConsoleLogger) Info(msg string) {
	fmt.Println("INFO: ", msg)
}

func (logSender LoggedSender) Send(ctx context.Context, to Email, body string) error {
	logSender.Log.Info("Envíado vía: " + logSender.Channel())
	return logSender.Sender.Send(ctx, to, body)
}
