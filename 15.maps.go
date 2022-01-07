package main

import "fmt"

func main() {

	// Los maps son listas que almacenan una cantidad de datos con clave y valor de la

	map1 := map[string]int{ // Se difine el map con el tipo de datos
		"key1": 1, // Se le asigna un valor a la clave
		"key2": 2,
		"key3": 3,
	}

	fmt.Println(map1) // map[key1:1 key2:2 key3:3]

	// Se puede agregar una nueva clave
	map1["key4"] = 4
	map1["key5"] = 5

	fmt.Println(map1) // map[key1:1 key2:2 key3:3 key4:4 key5:5]

	// Se puede acceder al valor de la clave
	fmt.Println(map1["key1"]) // 1

	// Se puede eliminar una clave
	delete(map1, "key1")
	fmt.Println(map1) // map[key2:2 key3:3 key4:4 key5:5

	// Inicializar un mapa con make, creara el map con claves vacias
	map2 := make(map[int]string) // Clave es un entero y valor tipo string
	map2[1] = "uno"
	map2[2] = "dos"
	map2[3] = "tres"
	fmt.Println(map2) // map[1:uno 2:dos 3:tres]

	// Se puede acceder a un valor de un mapa con una clave que no existe
	fmt.Println(map2[4]) // <nil> vacio

	// Se puede verificar si una clave existe en un mapa
	_, ok := map2[4]
	fmt.Println(ok) // false

	// Concatenar un mapa con otro
	map3 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}

	map4 := map[string]int{
		"key4": 4,
		"key5": 5,
		"key6": 6,
	}

	map5 := map3
	map5["key7"] = 7
	map5["key8"] = 8
	map5["key9"] = 9
	map5["key10"] = 10

	fmt.Println(map5) // map[key1:1 key10:10 key2:2 key3:3 key7:7 key8:8 key9:9]

	// Concatenar dos mapas
	map6 := map3
	for k, v := range map4 {
		map6[k] = v
	}

	fmt.Println(map6) // map[key1:1 key10:10 key2:2 key3:3 key4:4 key5:5 key6:6 key7:7 key8:8 key9:9]

}
