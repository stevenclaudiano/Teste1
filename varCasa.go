package main

import "fmt"

func main() {
	// Declaração das variáveis
	var operacao1, operacao2 string

	// Captura o estilo da casa
	fmt.Print("Digite o Estilo da casa que deseja construir: \n")
	fmt.Print("Casa de Madeira (1) ou Casa de Alvenaria (2) ou de Tijolo (3): \n")
	fmt.Scan(&operacao1)

	// Captura a cor da casa
	fmt.Print("Digite a cor da casa: \n")
	fmt.Print("Vermelha (1), Azul (2) ou Verde (3): \n")
	fmt.Scan(&operacao2)

	fmt.Println("\nRegistrando a casa...")
	registracasa(operacao1, operacao2)
}

func registracasa(operacao1, operacao2 string) {
	// Função auxiliar para padronizar a entrada
	//estilo := normalizarEntrada(operacao1)
	//cor := normalizarEntrada(operacao2)

	estilo := operacao1
	cor := operacao2

	// Switch que cobre todas as 9 combinações
	switch {
	// --- ESTILO: MADEIRA (1) ---
	case estilo == "1" && cor == "1":
		fmt.Println("Casa de Madeira Vermelha registrada com sucesso!")
	case estilo == "1" && cor == "2":
		fmt.Println("Casa de Madeira Azul registrada com sucesso!")
	case estilo == "1" && cor == "3":
		fmt.Println("Casa de Madeira Verde registrada com sucesso!")

	// --- ESTILO: ALVENARIA (2) ---
	case estilo == "2" && cor == "1":
		fmt.Println("Casa de Alvenaria Vermelha registrada com sucesso!")
	case estilo == "2" && cor == "2":
		fmt.Println("Casa de Alvenaria Azul registrada com sucesso!")
	case estilo == "2" && cor == "3":
		fmt.Println("Casa de Alvenaria Verde registrada com sucesso!")

	// --- ESTILO: TIJOLO (3) ---
	case estilo == "3" && cor == "1":
		fmt.Println("Casa de Tijolo Vermelha registrada com sucesso!")
	case estilo == "3" && cor == "2":
		fmt.Println("Casa de Tijolo Azul registrada com sucesso!")
	case estilo == "3" && cor == "3":
		fmt.Println("Casa de Tijolo Verde registrada com sucesso!")

	default:
		fmt.Println("Combinação inválida. Por favor, digite apenas os números (1, 2 ou 3) para Estilo e Cor.")
	}
}

/*
// Função auxiliar para converter o nome por extenso de volta para o número, se digitado.
func normalizarEntrada(entrada string) string {
	switch entrada {
	case "Madeira", "madeira":
		return "1"
	case "Alvenaria", "alvenaria":
		return "2"
	case "Tijolo", "tijolo":
		return "3"
	case "Vermelha", "vermelha":
		return "1"
	case "Azul", "azul":
		return "2"
	case "Verde", "verde":
		return "3"
	default:
		return entrada // Retorna o número (string) se foi digitado um número
	}
}

*/
