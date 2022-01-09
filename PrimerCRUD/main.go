package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"      // Para los logs
	"net/http" // Para la web
	"strconv"  // Para convertir string a int

	"github.com/gorilla/mux" // Para el enrutador
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

// API de tareas por lo que se definiran el tipo de datos
type task struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

// Todas las tareas
type allTask []task

var tasks = allTask{
	{
		Id:      1,
		Name:    "Tarea 1",
		Content: "Contenido de la tarea 1",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo Mi Primer API")
}

// Obtener todas las task
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Enviamos todas las tareas
	json.NewEncoder(w).Encode(tasks)
}

// Crear task
func createTask(w http.ResponseWriter, r *http.Request) {
	// Recibir info
	w.Header().Set("Content-Type", "application/json")
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al leer el body")
	}
	//Asignar info
	json.Unmarshal(reqBody, &newTask)
	newTask.Id = len(tasks) + 1 //AUmenten los id por default
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json") //EL tipo de contenido que se esta enviando al cliente
	w.WriteHeader(http.StatusCreated)                  // Codigo de creado
	json.NewEncoder(w).Encode(newTask)                 // Enviar la nueva tarea

}

//Obtenemos una task por id
func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}

	for _, task := range tasks {
		if task.Id == taskId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

// Borrar task
func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid User ID")
		return
	}

	for i, t := range tasks {
		if t.Id == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "The task with ID %v has been remove successfully", taskId)
		}
	}
}

// Actualizar task
func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	var updatedTask task

	if err != nil {
		fmt.Fprintf(w, "Invalid Id")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please Enter Valid Data")
	}
	json.Unmarshal(reqBody, &updatedTask)

	for i, t := range tasks {
		if t.Id == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)

			updatedTask.Id = t.Id
			tasks = append(tasks, updatedTask)

			// w.Header().Set("Content-Type", "application/json")
			// json.NewEncoder(w).Encode(updatedTask)
			fmt.Fprintf(w, "The task with Id %v has been updated successfully", taskId)
		}
	}

}

func main() {

	// Manejador de rutas
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Inicio)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET") // Para pedir por id
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3003", router))

}

func Inicio(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola Mar!")
	plantillas.ExecuteTemplate(w, "inicio", nil)
}
