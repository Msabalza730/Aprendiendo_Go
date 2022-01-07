package main

import "fmt"

func main() {
	// Aca se deben llamar las funciones definidas nuevas

	// Funcion de saludo
	saludo()

	// Funcion dato
	nombre, edad := dato()
	fmt.Println(nombre, edad)

	// Funcion suma
	var a, b int
	fmt.Print("Ingresa el numero a: ")
	fmt.Scanln(&a)
	fmt.Print("Ingresa el numero b: ")
	fmt.Scanln(&b)

	resultado := suma(a, b)
	fmt.Println("La suma de los numeros es: ", resultado)
}

// Definir una funcion
func saludo() {
	fmt.Println("Hola, soy una funcion ¿Cual es tu Nombre?")

	var nombre string
	fmt.Println("Ingresa tu nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("Hola", nombre, "mucho gusto en conocerte")
}

// Funcion con retorno de datos
func dato() (string, int) { // esta funcion tendra dos tipos de datos string e int

	var (
		nombre string
		edad   int
	)

	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanln(&nombre)

	fmt.Print("Ingresa tu edad: ")
	fmt.Scanln(&edad)

	return nombre, edad
}

func suma(a int, b int) int {
	suma := a + b
	return suma
}
