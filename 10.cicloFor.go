package main

import "fmt"

func main() {

	// En GO no existen el While asi que el FOR puede usarse como un bucle infinito
	// Ejemplo
	c := 0

	for c <= 5 {
		// Ciclo infinito
		//fmt.Println("Este bucle no se va a detener: "c)

		// Ciclo finito
		fmt.Println(c)
		c++ // Va a contar hasta 5
	}

	// Otra forma de hacerlo simplificado en 2 lineas
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// Agregando un BREAK
	for i := 0; i <= 15; i++ {
		if i == 5 {
			fmt.Println("Se detuvo el ciclo")
			break
		}
		fmt.Println(i)
	}

	// Agregando un CONTINUE
	for i := 0; i <= 15; i++ {
		if i == 5 {
			fmt.Println("Ejecutando.....")
			continue
			fmt.Println("Esto no se va a imprimir")
		}
		fmt.Println(i)
	}

}
