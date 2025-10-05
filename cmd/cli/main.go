// cmd/cli/main.go

package main

import (
	"analizador/lexico/lexer"
	"fmt"
)

func main() {
	codigoFuente := `
-- Prueba final del motor al 100%
entrada >>> [25/12/2025]
salida <<< variable_hora
`
	fmt.Println("--- Iniciando Análisis Léxico ---")
	fmt.Println("Código a analizar:")
	fmt.Println(codigoFuente)
	fmt.Println("---------------------------------")
	fmt.Println("Tokens encontrados:")

	// Creamos una nueva instancia de nuestro analizador.
	analizador := lexer.NuevoAnalizador(codigoFuente)

	// Bucle para obtener todos los tokens hasta el final del archivo.
	for {
		token := analizador.SiguienteToken()

		// Imprimimos el token encontrado en un formato legible.
		fmt.Printf("Línea: %d, Columna: %d \t| Tipo: %-20s \t| Lexema: '%s'\n",
			token.Linea, token.Columna, token.Type, token.Lexema)

		// Si encontramos el token de Fin de Archivo, terminamos el bucle.
		if token.Type == lexer.TOKEN_EOF {
			break
		}
	}

	fmt.Println("--- Fin del Análisis ---")
}
