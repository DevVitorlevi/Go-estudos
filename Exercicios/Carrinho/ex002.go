package main

import "fmt"

type Produto struct {
	ID      int
	Nome    string
	Preco   float32
	Estoque int
}

type CartItem struct {
	Produto
	Quantidade int
}

func CalcularCarrinho(Item CartItem) float64 {
	return float64(Item.Preco) * float64(Item.Quantidade)
}

func main() {
	Prod := Produto{ID: 1, Nome: "Notebook Gamer", Preco: 2800.00, Estoque: 10}
	Cart := CartItem{Produto: Prod, Quantidade: 2}

	Total := CalcularCarrinho(Cart)

	fmt.Printf("Produto: %s\n", Cart.Nome)
	fmt.Printf("Preço unitário: R$ %.2f\n", Cart.Preco)
	fmt.Printf("Quantidade: %d\n", Cart.Quantidade)
	fmt.Println("--------------------")
	fmt.Printf("Valor total: R$ %.2f\n", Total)
}
