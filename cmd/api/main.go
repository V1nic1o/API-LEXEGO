// cmd/api/main.go

package main

import (
	"analizador/lexico/api" // Importamos el paquete 'api' que contiene nuestro handler
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 1. Registramos nuestro manejador (handler) en una ruta específica.
	// Cada vez que alguien haga una petición a "/analizar", se ejecutará la función AnalizarHandler.
	http.HandleFunc("/analizar", api.AnalizarHandler)

	// 2. Definimos el puerto en el que escuchará nuestro servidor.
	puerto := ":8080"
	fmt.Printf("Servidor API iniciado. Escuchando en http://localhost%s\n", puerto)
	fmt.Println("Esperando peticiones en el endpoint /analizar ...")

	// 3. Iniciamos el servidor.
	// http.ListenAndServe se queda escuchando por peticiones y nunca termina,
	// a menos que haya un error fatal.
	log.Fatal(http.ListenAndServe(puerto, nil))
}
