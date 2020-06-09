package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calculadora struct{}

func (calculadora) operate(entrada string, operador string) int {

	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])

	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		return operador1 + operador2
	case "-":
		fmt.Println(operador1 - operador2)
		return operador1 - operador2
	case "*":
		fmt.Println(operador1 * operador2)
		return operador1 * operador2
	case "/":
		fmt.Println(operador1 / operador2)
		return operador1 / operador2
	default:
		fmt.Println("Operador no soportado")
		return 0
	}
}

func parsear(valores string) int {
	operador, _ := strconv.Atoi(valores)
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	entrada := leerEntrada()
	operador := leerEntrada()

	c := calculadora{}
	c.operate()

}
