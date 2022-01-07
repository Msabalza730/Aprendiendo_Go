package main

import "fmt"

func main() {

	array1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // array definido con una lista de enteros

	// iterar un array con for
	for _, num := range array1 {
		fmt.Println(num)
	}

	// iterar un array con for con un indice
	for i, num := range array1 {
		fmt.Println(i, ":", num)
	}

	// sumar valores de array
	suma := 0 //inicializando en cero
	for _, num := range array1 {
		suma += num
	}
	fmt.Println("La suma de los valores es: ", suma)

	// iterar un mapa con for y range
	map1 := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	for k, v := range map1 {
		fmt.Println(k, ":", v)
	}

}
