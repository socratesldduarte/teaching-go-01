package main

import "fmt"

func main() {
	/*
		Essa estrutura abaixo (for sem nenhum parâmetro) é um loop infinito
		  que será executado em loop indefinidamente, até encontrar um ponto de saída
		  (no nosso caso, um return na opção 9, que encerra a execução do programa).
		Nesse nosso exemplo, está sendo utilizado para ficar executando até que
		  seja selecionada a opção de saída do menu
	*/
	for {
		// APRESENTAÇÃO DO MENU
		fmt.Println("*** CONVERSÃO DE TEMPERATURA ***\nEscolha uma opção do menu")
		fmt.Println("1 - Converter Celsius para Fahrenheit")
		fmt.Println("2 - Converter Fahrenheit para Celsius")
		fmt.Println("9 - Sair do programa")

		// LENDO A OPÇÃO SELECIONADA
		var option int
		fmt.Scan(&option)

		// O QUE FAZER PARA CADA OPÇÃO SELECIONADA
		switch option {
		case 1:
			fmt.Println("Informe a temperatura Celsius")
			var celsius int
			fmt.Scan(&celsius)
			// Calculo de Fahrenheit
			fahrenheit := float32(celsius)*1.8 + 32
			fmt.Printf("A conversão de %d Celsius para Fahrenheit resulta em %.2f\n", celsius, fahrenheit)
		case 2:
			fmt.Println("Informe a temperatura Fahrenheit")
			var fahrenheit int
			fmt.Scan(&fahrenheit)
			// Calculo de Celsius
			celsius := (float32(fahrenheit)*10 - 320) / 18
			fmt.Printf("A conversão de %d Fahrenheit para Celsius resulta em %.2f\n", fahrenheit, celsius)
		case 9:
			fmt.Println("ENCERRANDO A EXECUÇÃO...")
			return
		default:
			fmt.Println("OPÇÃO INVÁLIDA")
		}
	}
}
