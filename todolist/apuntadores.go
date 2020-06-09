package main2

import "fmt"

func main2() {
	x := 25
	fmt.Println(x)
	fmt.Println("Direccion de memoria inicial")
	fmt.Println(&x)
	cambiarValor(x)

	compararVariables()

}

func compararVariables() {
	variable1 := 10
	fmt.Println(variable1)
	variable2 := variable1
	fmt.Println(variable2)

	variable1 = 20
	fmt.Println(variable1)

}

func cambiarValor(a int) {
	fmt.Println(&a)
	a = 36
}
