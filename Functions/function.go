package main

import "fmt"

func Somar(n1, n2 int) int {
	return n1 + n2
}

func SomaeSubtrair(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtrair := n1 - n2
	return soma, subtrair
}
func main() {
	fmt.Println(Somar(10, 2))
	// rSoma, rSubtrair := SomaeSubtrair(20, 10)
	// fmt.Println(rSoma, rSubtrair)

	rSoma, _ := SomaeSubtrair(20, 10)
	fmt.Println(rSoma)
}
