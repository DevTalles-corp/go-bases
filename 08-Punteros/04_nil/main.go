package main

import "fmt"

type User struct {
	Name string
}

type Config struct {
	Host string
	Port int
}

func printName(user *User) {
	if user == nil {
		fmt.Println("User está vacío (nil)")
		return
	}
	fmt.Println("User: ", user.Name)
}

func NewConfig() *Config {
	return &Config{Host: "localhost", Port: 8080}
}

func main() {
	var user *User // nil
	printName(user)

	// user = &User{Name: "Ricardo"}
	// printName(user)

	my_config := NewConfig()
	fmt.Println("configuración: ", my_config.Host, my_config.Port)
}
