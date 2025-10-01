package main

import "fmt"

func main() {

	exibintro()
	comando := comando()
	escolha(comando)

}

func exibintro() {
	nome := "Steven" // Atribuição de variavel.
	Cidade := "Itapui"

	fmt.Println("O meu nome é:", nome)
	fmt.Println("O seu Templo Superior é:", Cidade)

	fmt.Println("-----------------------------------------")
	fmt.Println("ESCOLHA O SEU PERFIL \n")
	fmt.Println("Perfil 1 Lutador \n")
	fmt.Println("Perfil 2 Ninja \n")
	fmt.Println("Perfil 3 Mago \n")
	fmt.Println("Perfil 4 Tanque \n")

}

func comando() int {
	var comandolido int
	fmt.Scan(&comandolido)
	fmt.Println("O Perfil Escolhido foi: ", comandolido)

	return comandolido

}

func escolha(comando int) {
	nome := "Steven"
	switch comando {
	case 1:
		fmt.Println(nome, ",O seu Perfil é de um Eximio Lutador Mestre!")
	case 2:
		fmt.Println(nome, ",O seu Perfil é de um Eximio Ninja Das Sombras")
	case 3:
		fmt.Println(nome, ",O seu Perfil é de um Eximio Mago Superior")
	case 4:
		fmt.Println(nome, ",O seu Perfil é de um Eximio Tanque de Vida")
	default:
		fmt.Println(nome, ",Fim")
		break
	}
}
