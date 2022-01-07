package main

import "fmt"

// Crear un tipo de datos

type dato int
type dato2 string

func main() {

	var edad dato // asignando el tipo de dato creado
	var nombre dato2

	edad = 20
	nombre = "Juan"
	fmt.Println(nombre, "tiene", edad, "años")

	// Llamando la Estructura
	SuperHeroe1 := SuperHeroe{

		nombre: "Maryori",
		edad:   28,
		poder:  []string{"Volar", "Programar"},
	}
	fmt.Println(SuperHeroe1) //{Maryori 28 [Volar Programar]}

	// Se pueden crear varias estructuras
	SuperHeroe2 := SuperHeroe{

		nombre: "Filomena",
		edad:   6,
		poder:  []string{"Nadar", "Comer y beber"},
	}
	fmt.Println(SuperHeroe2) //{Filomena 6 [Nadar Comer y beber]}
}

// En GO no hay PO pero se usan estructuras
type SuperHeroe struct {
	nombre string
	edad   int
	poder  []string
}
