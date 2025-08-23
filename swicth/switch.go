package main

import "fmt"

func diasDaSemana(dia int) string {
	switch dia {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terca-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabado"
	default:
		return "Numero Invalido"
	}

}
func main() {
	fmt.Println(diasDaSemana(1))
}
