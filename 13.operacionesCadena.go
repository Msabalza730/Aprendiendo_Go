package main

// Importando paquetes
import (
	"fmt"
	"strings" // Importar el paquete strings  para realizar operaciones con cadena
)

func main() {
	// Se verifica dentro del paquete que contiene "Contains"
	fmt.Println(strings.Contains("seafood", "foo")) // true

	// Para contar se usa Count
	fmt.Println(strings.Count("cheese", "e")) // 3

	// variable string para convertirla en mayusculas o minusculas
	txt := "Hola Maryori, esto es un string"
	fmt.Println(strings.ToUpper(txt)) // HOLA MARYORI, ESTO ES UN STRING
	fmt.Println(strings.ToLower(txt)) // hola maryori, esto es un string

	//Para Reemplazar
	fmt.Println(strings.Replace(txt, "Maryori", "Mar", 1))              // HOLA Mar, esto es un string, el 1 es para decirle cuanto vamos a reemplazar
	fmt.Println(strings.ReplaceAll(txt, "Hola Maryori", "Hello World")) // ReplaceAll reemplaza todo lo que coincida con el primer parametro

	// Split
	fmt.Println(strings.Split(txt, "Maryori")) // HOLA , esto es un string --> Estor es un array

	// Convirtiendo a un slice
	var slicen1 []string

	slicen1 = append(strings.Split(txt, " "))
	fmt.Println(slicen1, len(slicen1)) // [Hola Maryori, esto es un string] .. tamaño 6

}
