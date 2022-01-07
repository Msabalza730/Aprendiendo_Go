package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	fmt.Println("Limpiando pantalla en 1 segundo......!")
	time.Sleep(1 * time.Second)
	CallClear()

	fmt.Print("****************  Calculadora de Mar ********************************\n")

	var a, b, opcion int
	fmt.Print("Ingresa el numero a: ")
	fmt.Scanln(&a)
	fmt.Print("Ingresa el numero b: ")
	fmt.Scanln(&b)
	fmt.Print("Opciones disponibles... \n")
	fmt.Print("Suma => 1 \n")
	fmt.Print("Resta => 2 \n")
	fmt.Print("Multiplicacion => 3 \n")
	fmt.Print("Division => 4 \n")
	fmt.Print("Salir => 5 \n")
	fmt.Print("Ingrese que operacion desea hacer: \n")
	fmt.Scanln(&opcion)

	switch opcion {
	case 1:
		fmt.Print("La suma es: ", Sumar(a, b))
	case 2:
		fmt.Print("La resta es: ", Restar(a, b))
	case 3:
		fmt.Print("La multiplicacion es: ", Multiplicacion(a, b))
	case 4:
		fmt.Print("La division es: ", Division(a, b))
	case 5:
		fmt.Print("Gracias por usar la calculadora de Mar!")

	}
}

func Sumar(a int, b int) int {
	suma := a + b
	return suma
}

func Restar(a int, b int) int {
	resta := a - b
	return resta
}

func Multiplicacion(a int, b int) int {
	multi := a * b
	return multi
}

func Division(a int, b int) int {
	div := a / b
	return div
}
