package main

import "fmt"

func main() {

	nome, idade := retornaNome()
	fmt.Println("Meu nome é", nome, "A idade é:", idade)

}

func retornaNome() (string, int) {
	nome := "Steven"
	idade := 30
	return nome, idade
}

// func main() {

// 	nome, _ := retornaNome()
// 	fmt.Println("Meu nome é", nome, "A idade é:")		COLOCAR O _ PARA ESCOLHER APENAS 1 PASSADA DE PARAMETRO

// }
