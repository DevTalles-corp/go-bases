package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Email string

// Money en centavos
type Money int64

var emailRe = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

func NewEmail(v string) (Email, error) {
	v = strings.TrimSpace(strings.ToLower(v))
	if !emailRe.MatchString(v) {
		return "", fmt.Errorf("email inválido: %q", v)
	}
	return Email(v), nil
}

func NewMoneyFromCents(cents int64) (Money, error) {
	if cents < 0 {
		return 0, fmt.Errorf("money no puede ser negativo")
	}
	return Money(cents), nil
}

func (m Money) String() string {
	d := int64(m) / 100
	c := int64(m) % 100
	return fmt.Sprintf("$%d.%02d", d, c)
}
