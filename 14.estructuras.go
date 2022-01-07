// Las funciones pueden ser publicas o privadas y esto depende la estructura que tengan

package main

import "fmt"

func main() {
	// Funcion suma
	var a, b int
	fmt.Print("Ingresa el numero a: ")
	fmt.Scanln(&a)
	fmt.Print("Ingresa el numero b: ")
	fmt.Scanln(&b)

	resultado := suma(a, b)
	fmt.Println("La suma de los numeros es: ", resultado)

	resultado1 := Resta(a, b)
	fmt.Println("La resta de los numeros es: ", resultado1)

}

// Estructura Privada
/*  minúscula (no exportar) como primer caracter
type miEstructura type {
	miMetodo()}
*/
func suma(a int, b int) int {
	suma := a + b
	return suma
}

// Estructura Publica
/*  mayúscula (exportar)  como primer caracter
type MiEstructura type {
	miMetodo()}
*/
func Resta(a int, b int) int {
	Resta := a - b
	return Resta
}
