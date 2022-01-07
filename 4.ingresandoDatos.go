package main

import "fmt"

func main() {

	var (
		nombre string
		edad   int
		pi     float64
	)

	fmt.Print("************ Bienvenido ********")
	//fmt.Scanf("%s", &nombre) // Para indicar que tipo de dato vamos a recibir

	fmt.Print("\nIngrese su Nombre: ")
	fmt.Scan(&nombre) // Para indicar que tipo de dato vamos a recibir

	fmt.Print("\nIngrese su Edad: ")
	fmt.Scan(&edad)

	fmt.Print("\nIngrese su PI: ")
	fmt.Scan(&pi)

	fmt.Print("************ Adios ********")
}
