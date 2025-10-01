package main

import "fmt"

type scarros struct {
	Nome   string
	Marca  string
	Modelo string
	Ano    int
	Preco  float64
	Usado  bool
}

func main() {

	fmt.Println(cadastroCarro())

}

func cadastroCarro() []scarros {
	var carros []scarros

	carros = append(carros, scarros{
		Nome:  "Mustang",
		Marca: "Ford",
		Ano:   2011,
		Preco: 261.000,
		Usado: false,
	})

	carros = append(carros, scarros{
		Nome:  "Civic",
		Marca: "Honda",
		Ano:   2019,
		Preco: 150000.00,
		Usado: true,
	})

	carros = append(carros, scarros{
		Nome:  "Corolla",
		Marca: "Toyota",
		Ano:   2021,
		Preco: 180000.00,
		Usado: false,
	})
	return carros

}

func exibircarros(c scarros) {
	fmt.Println("NOME: ", c.Nome)

}
