package main

import "fmt"

type Address struct {
	City  string
	State string
}

type Profile struct {
	Bio string
}

type User struct {
	Name    string
	Addr    Address
	Profile *Profile //Opcional
}

func main() {
	user := User{
		Name: "Ana",
		Addr: Address{
			City:  "Buenos Aires",
			State: "Mar de la plata",
		},
	}

	if user.Profile == nil {
		fmt.Println("SIN PERFIL")
	}

	user.Profile = &Profile{Bio: "Soy Ricardo Cuéllar y me gusta programar en go"}

	fmt.Println("Bio: ", user.Profile.Bio)
	fmt.Println(user)

}
