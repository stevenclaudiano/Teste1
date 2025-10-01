package main

import "fmt"

type Tcarros1 struct { // Tcarros1 é o nome da Strutura
	Nome   string
	Marca  string
	Modelo string
	Ano    int
	Preco  float64
	Usado  bool
}

func main() {
	// 1. Chama a função que retorna TODOS os carros1 (um slice)
	carros1 := cadastrocarros1()

	// 2. Loop para exibir cada carro individualmente
	for _, i := range carros1 {
		exibircarros1(i)
	}
}

func cadastrocarros1() []Tcarros1 {
	var carros1 []Tcarros1

	carros1 = append(carros1, Tcarros1{
		Nome:  "Mustang",
		Marca: "Ford",
		Ano:   2011,
		Preco: 261.000,
		Usado: false,
	})

	carros1 = append(carros1, Tcarros1{
		Nome:  "Civic",
		Marca: "Honda",
		Ano:   2019,
		Preco: 150000.00,
		Usado: true,
	})

	carros1 = append(carros1, Tcarros1{
		Nome:  "Corolla",
		Marca: "Toyota",
		Ano:   2021,
		Preco: 180000.00,
		Usado: false,
	})
	return carros1 // retorna a variavel contendo as informações.

}

func exibircarros1(c Tcarros1) {

	fmt.Println("Nome: ", c.Nome)
	fmt.Println("Marca: ", c.Marca)
	fmt.Println("Ano: ", c.Ano)
	fmt.Println("Preco: ", c.Preco)
	fmt.Println("Usado: ", c.Usado)
	fmt.Println("-----------------------")
}
