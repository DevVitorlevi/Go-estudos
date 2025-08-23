package main

import "fmt"

type Usuario struct {
	id    int
	nome  string
	email string
	ativo bool
}

func main() {
	newUser := Usuario{1, "Vitor", "vitor@gmail.com", true}
	fmt.Printf("%+v\n", newUser)
}
