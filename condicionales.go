package main

import "fmt"

func main() {

	if false {
		fmt.Println("Hola se cumple una condicion")
	} else {
		fmt.Println("Hola se cumple otra condicion")
	}

	// Crea un sistema que ingrese un numero entero y diga si es par o impar

	var a int

	// Entrada de datos
	fmt.Print("Ingrese el numero: ")
	fmt.Scan(&a)

	// Proceso
	if a%2 == 0 {
		fmt.Println("El numero es par")
	} else {
		fmt.Println("El numero es impar")
	}

}
