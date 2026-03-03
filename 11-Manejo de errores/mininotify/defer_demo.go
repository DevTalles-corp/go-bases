package main

import (
	"fmt"
	"os"
)

func WriteReport(path string, content string) (err error) {
	file, err := os.Create(path)

	if err != nil {
		return fmt.Errorf("Archivo creado: %q: %w", path, err)
	}

	defer func() {
		closeError := file.Close()
		if closeError != nil && err == nil {
			err = fmt.Errorf("Al cerrar el archivo: %q: %w", path, closeError)
		}
	}()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("Al escribir en el archivo. %q: %w", path, err)
	}

	return nil

}

func ShadowingError() error {
	var err error //nil

	if _, err := os.Stat("no-existe.txt"); err != nil {
		_ = err
	}

	return err
}
