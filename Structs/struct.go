package main

import "fmt"

type user struct {
	id         int
	name       string
	idade      int
	disponivel bool
	empresa    empresa
}

type empresa struct {
	name  string
	local string
}

func main() {
	empresa := empresa{"Golden", "Fortaleza-CE"}
	User := user{1, "Vitor", 10, true, empresa}
	fmt.Println(User)
}
