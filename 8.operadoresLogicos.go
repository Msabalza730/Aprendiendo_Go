package main

import "fmt"

// Not !

// AND  &&

// OR ||

func main() {

	fmt.Println("Operadores logicos")
	fmt.Println(!true)
	fmt.Println(false && true)
	fmt.Println(true || true)

	fmt.Println(2 == 2 && 3 != 4)

	// Realizar un sistema que pida ingresar un numero de 1a5 y devuelva  el numero escrito

	var (
		n   int
		out string
	)

	fmt.Println("Ingrese un numero del 1 al 5:  ")
	fmt.Scanln(&n)

	if n > 5 {
		fmt.Println("Numero invalido")
	} else {
		switch n {
		case 1:
			out = "Uno"
		case 2:
			out = "Dos"
		case 3:
			out = "Tres"
		case 4:
			out = "Cuatro"
		case 5:
			out = "Cinco"
		}
		fmt.Println(out)
	}

}
