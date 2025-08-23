package main

import "fmt"

func main() {
	var variavel3 int
	var ponterio *int

	variavel3 = 100
	ponterio = &variavel3

	//Ponteiro e uma referencia de memoria

	fmt.Println(variavel3, ponterio)  // Local da Variavel na Memoria
	fmt.Println(variavel3, *ponterio) // Valor da Varialvel na Memoria
}
