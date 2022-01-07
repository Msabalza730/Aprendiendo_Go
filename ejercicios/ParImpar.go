package main

import "fmt"

func main() {

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
