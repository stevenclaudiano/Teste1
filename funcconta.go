package main

import "fmt"

func somar(teste string, n1 int, n2 int) (int, string) {
	return n1 + n2, teste
}

func main() {
	nome, soma1 := somar("Steven", 40, 10)
	fmt.Println(nome, soma1)
}
