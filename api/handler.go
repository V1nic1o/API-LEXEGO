// api/handler.go

package api

import (
	"analizador/lexico/lexer"
	"encoding/json"
	"net/http"
)

// SolicitudAnalisis (sin cambios)
type SolicitudAnalisis struct {
	CodigoFuente string `json:"codigoFuente"`
}

// setupCORS es una nueva función para configurar las cabeceras de CORS.
func setupCORS(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// AnalizarHandler (actualizado para manejar OPTIONS)
func AnalizarHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Configuramos las cabeceras de CORS para CADA petición.
	setupCORS(&w, r)

	// 2. Manejamos la petición de sondeo (preflight) del navegador.
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// 3. Verificamos que el método sea POST (la lógica original).
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// El resto de la función sigue exactamente igual que antes.
	var solicitud SolicitudAnalisis
	err := json.NewDecoder(r.Body).Decode(&solicitud)
	if err != nil {
		http.Error(w, "Error al decodificar el JSON de la solicitud", http.StatusBadRequest)
		return
	}

	analizador := lexer.NuevoAnalizador(solicitud.CodigoFuente)
	var tokens []lexer.Token
	for {
		token := analizador.SiguienteToken()
		tokens = append(tokens, token)
		if token.Type == lexer.TOKEN_EOF {
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tokens)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta JSON", http.StatusInternalServerError)
	}
}
