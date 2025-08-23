package main

import "fmt"

type person struct {
	name   string
	age    int
	height float64
}

type student struct {
	person
	course string
	school string
}

func main() {
	p1 := person{"Vitor", 17, 1.71}
	st1 := student{p1, "Desenvolvimento de Sistemas ", "EEEP Jaime"}

	fmt.Println(st1)
}
