/// https://es.wikipedia.org/wiki/Go_(lenguaje_de_programaci%C3%B3n)

package main

import (
	"fmt"
	"log"
	"net/http"
)

func holaFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo!")
}

func main() {
	http.HandleFunc("/", holaFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Ir a http://localhost:8080/ y ver el resultado
