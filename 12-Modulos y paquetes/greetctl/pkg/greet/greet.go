package greet

import "strings"

func Hello(name string) string {
	name = normalizeName(name)
	return "Hola, " + name + " :)"
}

func normalizeName(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "mundo"
	}

	return name
}
