# Fundamentos de GO
## Tipos de Dados e Conversão com Estruturas de controle e repetição

Nessa primeira aula, apresentamos a estrutura de um programa em Go, incluindo a utilização de uma primeira biblioteca 
padrão (Standard library) chamada fmt.

```go
package main

import "fmt"
```
O arquivo principal (main.go) referencia o nome do pacote que está send criado (em nosso caso, o programa mian).
E utilizamos a instrução import fmt para "adicionar o pacote fmt" ao nosso código. Assim podemos utilizar os métodos de 
fmt.

Em nosso primeiro exercício, todo o código está compreendido na função main (essa é a função que é o "ponto de partida"
da execução do programa). Normalmente teremos outras funções em nosso programa, mas por fim didáticos, todo o código 
está na própria função main. Poderíamos, por exemplo, ter uma função que recebe o valor de uma temperatura em Celsius e 
retorna o valor em Fahrenheit e uma função que recebe um valor em Farenheit e converte em Celsius. Ou mesmo uma função 
mais genérica que recebe múltiplos valores (por exemplo, temperatura de origem, temperatura de destino e valor da 
temperatura medida) e converte a temperatura medida em "origem" para "destino".

Dentro da função main(), temos uma estrutura de repetição for, sem nenhum parâmetro. Essa estrutura causa um loop 
infinito (ou seja, esse trecho será executado repetidamente até que haja alguma condição de saída).
Então, para compreensão, esse trecho do programa (todo o conteúdo do for) será executado repetidas vezes até que ocorra 
alguma ação explícita para saída.

```go
for {
	// instruções
}
```

O motivo é ficar "repetindo" esse trecho até que o usuário indique que quer sair do programa.
Dentro do for, temos primeiro a exibição de um menu, feito com a instrução Println do pacote fmt.

```go
    // APRESENTAÇÃO DO MENU
    fmt.Println("*** CONVERSÃO DE TEMPERATURA ***\nEscolha uma opção do menu")
    fmt.Println("1 - Converter Celsius para Fahrenheit")
    fmt.Println("2 - Converter Fahrenheit para Celsius")
    fmt.Println("9 - Sair do programa")
```

O Println é uma instrução de "exiba o conteúdo a seguir e então salte uma linha".
Então, para cada linha acima, haverá um salto de linha após exibir o conteúdo. Observe que na primeira instrução de 
Println existe uma sequência "\n". Essa sequência provo o início de uma nova linha. Ou seja, essa primeira instrução 
ocupará DUAS linhas: a primeira com o título *** CONVERSÃO DE TEMPERATURA *** e a segunda com o texto Escolha uma opção 
do menu. O mesmo resultado poderia ser alcançado com dois Println distintos.

A seguir, temos o trecho para "ler" a opção do usuário (ou seja, esperar que o usuário digite a opção desejada).

```go
    // LENDO A OPÇÃO SELECIONADA
    var option int
    fmt.Scan(&option)
```

Nesse trecho temos a declaração de uma variável option, do tipo inteiro. Essa variável será responsável por obter a 
opção informada pelo usuário e comparar com as opções disponíveis, através de uma instrução de tomada de decisão 
chamada switch.

```go
		// O QUE FAZER PARA CADA OPÇÃO SELECIONADA
		switch option {
		case 1:
			// Calculo de Fahrenheit
		case 2:
			// Calculo de Celsius
		case 9:
			fmt.Println("ENCERRANDO A EXECUÇÃO...")
			return
		default:
			fmt.Println("OPÇÃO INVÁLIDA")
		}
```

A instrução switch substitui a estrutura if - else if quando diversas condições devem ser testadas sobre uma mesma 
variável. Em nosso exercício, devemos comparar o valor informado pelo usuário (option) com as opções de menu.
Basicamente, se o valor for 1, o programa executará o trecho em "case 1", se for 2 o trecho em "case 2" e assim por 
diante.
Se nenhuma das condições for satisfeita (por exemplo, se digitar 3, 4, 5 ou mesmo um valor não numérico) o trecho em 
"default" será executado.
E após a execução do trecho específico, o programa prosseguirá sua execução normalmente. Nesse caso, ao final do switch 
temos o final do for, o que significa que o programa executará novamente todo o trecho do for.
Porém é importante ressaltar que a opção  "9" possui um instrução "return", que provoca o encerramento da função atual 
(sendo a main, equivale a encerrar o programa).

E os trechos das opções 1 e 2 basicamente solicitam a entrada de um valor inteiro ao usuário (a temperatura Celsius ou 
Fahrenheit) conforme tenha sido a opção selecionada e converte e exibe o resultado.

```go
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
```

Observe que os trecho do case 1 e case 2 são bastante parecidos - e são candidatos a serem transformados em uma função 
que converte temperatura em uma melhoria desse código.