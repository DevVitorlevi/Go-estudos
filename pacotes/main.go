package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxi"
)

func main() {
	fmt.Println("Escrevendo dentro do pacote main")
	auxi.Escrever()
	error := checkmail.ValidateFormat("Vitor123@gmail.com")
	fmt.Println(error)
}
