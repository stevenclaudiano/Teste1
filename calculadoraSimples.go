package main

import "fmt"

func main() {
	var operacao int
	var num1, num2 float64
	var resultado float64
	for {
		fmt.Println("Escolha a operação:")
		fmt.Println("1 - Adição")
		fmt.Println("2 - Subtração")
		fmt.Println("3 - Multiplicação")
		fmt.Println("4 - Divisão")
		fmt.Println("5 - Sair")
		fmt.Scan(&operacao)

		if operacao == 5 {
			fmt.Println("Saindo...")
			break
		}

		fmt.Print("Digite o primeiro número: ")
		fmt.Scan(&num1)
		fmt.Print("Digite o segundo número: ")
		fmt.Scan(&num2)

		resultado = calcularoperacao(operacao, num1, num2)
		fmt.Printf("Resultado: %.2f\n", resultado)
	}
}

func calcularoperacao(operacao int, num1 float64, num2 float64) float64 {
	switch operacao {
	case 1:
		return num1 + num2
	case 2:
		return num1 - num2
	case 3:
		return num1 * num2
	case 4:
		if num2 != 0 {
			return num1 / num2
		} else {
			fmt.Println("Erro: Divisão por zero não é permitida.")
			return 0
		}
	case 5:
		fmt.Println("Saindo...")
		return 0
	default:
		fmt.Println("Operação inválida.")
		return 0
	}

}

// Atualizado
