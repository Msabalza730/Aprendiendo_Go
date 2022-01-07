package main

import "fmt"

func main() {

	//Definir Variables
	var nombre string

	nombre = "Maryori"

	fmt.Println(nombre)

	nombre = "Alicia"

	fmt.Println(nombre)

	//Definir Variables multiples

	var a, b int = 1, 2

	fmt.Println(a, b)

	var (
		pi     float64 = 3.14
		bol    bool    = false
		cadena         = "Texto 01" // Queda por default de tipo string
		edad           = 28
	)

	// Mostrando datos por pantalla con salto de linea
	println("***************** Aprendiendo Go *****************", "\n", "MI edad es: ", edad, "\n", "Tipo float: ", pi, "\n", "Tipo booleano: ", bol, "\n", "Cadena de caracteres: ", cadena)

	v := 23 // Se pone de tipo entero  por default con dos puntos
	println("Variable int por los ':=' ", v)

	// Contantes almacenan pero no se puede modificar su valor de la
	const n = 35

	println("Esto es una Constante: ", n)

	// PALABRAS RESERVADAS
	/*
		BREAK
		CASE
		CHAN
		CONST
		CONTINUE
		DEFAULT
		DEFER
		ELSE
		FALLTHROUGH
		FOR
		FUNCT
		GO
		GOTO
		IF
		IMPORT
		INTERFACE
		MAP
		PACKAGE
		RANGE
		RETURN
		SELECT
		STRUCT
		SWITCH
		TYPE
		VAR
	*/

}
