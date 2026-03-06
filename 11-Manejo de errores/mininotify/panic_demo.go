package main

import "fmt"

func RunSafely(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic recuperado: %v", r)
		}
	}()

	fn()
	return nil
}

func MustTemplate(name string) string {
	panic("No se encontro template: " + name)
}
