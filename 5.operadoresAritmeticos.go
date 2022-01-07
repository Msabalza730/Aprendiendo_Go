package main

import "fmt"

/*

	+  Sumar
	- Restar
/	* Multiplicar
	/ Dividir
	% Mod division
*/

func main() {

	var a, b int

	// Entrada de datos
	fmt.Print("Ingrese el primer numero: ")
	fmt.Scan(&a)

	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scan(&b)

	// Proceso
	suma := a + b

	resta := a - b

	multiplication := a * b

	division := a / b

	mod := a % b

	// Salida de datos
	fmt.Println("\nLa suma es: ", suma)
	fmt.Println("La resta es: ", resta)
	fmt.Println("La multiplicacion es: ", multiplication)
	fmt.Println("La division es: ", division)
	fmt.Println("El modulo es: ", mod)
}
