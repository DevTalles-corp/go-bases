package main

import (
	"context"
	"fmt"
)

type EmailSenderFake struct {
	Sent int
}

func (sender *EmailSenderFake) Send(ctx context.Context, to Email, body string) error {
	_ = ctx
	sender.Sent++
	fmt.Println("EMAIL SENT: ", sender.Sent, "TO: ", to)
	fmt.Println("BODY: ", body)
	return nil
}
