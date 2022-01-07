package main

import "fmt"

/*

	== Igualdad
/	!= Diferente
	> Mayor que
	< Menor que
	>= Mayor o igual que
	<= Menor o igual que
*/

func main() {

	var a, b int

	// Entrada de datos
	fmt.Print("Ingrese el primer numero: ")
	fmt.Scan(&a)

	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scan(&b)

	// Proceso

	// Igualdad
	fmt.Println("\nLos numeros son iguales: ", a == b)
	// Diferente
	fmt.Println("\nLos numeros son diferentes: ", a != b)
	// Mayor que
	fmt.Println("\nEl primer numero es mayor que el segundo: ", a > b)
	// Menor que
	fmt.Println("\nEl primer numero es menor que el segundo: ", a < b)
	// Mayor o igual que
	fmt.Println("\nEl primer numero es mayor o igual que el segundo: ", a >= b)
	// Menor o igual que
	fmt.Println("\nEl primer numero es menor o igual que el segundo: ", a <= b)

}
