// cmd/api/main.go
package main

import (
	"analizador/lexico/api" // Corregido: La ruta correcta al paquete 'api'
	"fmt"
	"log"
	"net/http"
	"os" // 1. IMPORTAMOS EL PAQUETE 'os' para leer variables de entorno
)

func main() {
	// 2. LEEMOS LA VARIABLE DE ENTORNO 'PORT' QUE NOS DARÁ RENDER
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Si no existe (en tu compu), usamos 8080 como respaldo
	}

	// Registramos nuestro manejador (handler) en la ruta "/analizar"
	http.HandleFunc("/analizar", api.AnalizarHandler)

	fmt.Printf("Servidor API iniciado. Escuchando en el puerto :%s\n", port)
	fmt.Println("Esperando peticiones en el endpoint /analizar ...")

	// 3. INICIAMOS EL SERVIDOR USANDO EL PUERTO CORRECTO
	// Render necesita que escuchemos en "0.0.0.0:<puerto>" que se abrevia como ":<puerto>"
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
