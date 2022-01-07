package main

import "fmt"

func main() {

	// Llamando la Estructura
	SuperHeroe1 := SuperHeroe{

		nombre: "Maryori",
		edad:   28,
		poder:  []string{"Volar", "Programar"},
	}
	fmt.Println(SuperHeroe1) //{Maryori 28 [Volar Programar]}
	SuperHeroe1.NuevoSuper("Mayra")

	// Llamando la Estructura nueva -- Villano
	Villano1 := new(Villano)
	Villano1.nombreVillano = "Aracnido"
	Villano1.edadVillano = 38
	Villano1.poderVillano = []string{"Veneno", "Venganza"}
	Villano1.NuevoSuper("Aracnido")
	fmt.Println(Villano1) //{Aracnido 38 {Maryori 28 [Volar Programar]}}
}

// En GO no hay PO pero se usan estructuras
type SuperHeroe struct {
	nombre string
	edad   int
	poder  []string
}

// Crear Metodo
// NuevoHeroe es un metodo de la esctructura SuperHeroe
func (metodo SuperHeroe) NuevoSuper(nombre string) {
	fmt.Printf("Nuevo SuperHeroe: %s", nombre, metodo.nombre)
}

// Otra estructura que hereda de la que existe
type Villano struct {
	nombreVillano string
	edadVillano   int
	poderVillano  []string
	SuperHeroe
}
