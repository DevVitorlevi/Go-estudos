package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//Iterar Array/Slices

	nomes := []string{"Vitor", "Caio", "Jota"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}
}
