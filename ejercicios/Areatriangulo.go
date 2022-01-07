package main

import "fmt"

func main() {

	var base, altura int
	fmt.Print("Ingrese la base del triangulo: ")
	fmt.Scanln(&base)
	fmt.Print("Ingrese la altura del triangulo: ")
	fmt.Scanln(&altura)
	fmt.Print("El area del triangulo es: ", base*altura, "\n")
}
