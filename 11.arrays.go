package main

import "fmt"

func main() {
	// Array es un almacenador de datos fijos y de tamaño fijo

	//Array vacio
	var array1 [2]string
	fmt.Println(array1)

	// Array con valores
	array1[0] = "Hola"
	array1[1] = "Mundo"
	fmt.Println(array1)

	// Slicen es como un array pero almacenan datos indeterminados sin tamaño fijo
	// slicen vacio

	var slicen0 []string

	// Agregando valores a un slice
	slicen0 = append(slicen0, "Hola", "Mundo", "Mundo", "Mundo", "Mundo", "Mundo", "Mar", "Mar", "Mar", "Mar", "Mar", "Mar", "")
	//slicen0 = append(slicen0, "Hola2", "M", "M", "M", "Mar", "Mar", "Mar", "") // se agregan al mismo slice
	fmt.Println(slicen0)

}
