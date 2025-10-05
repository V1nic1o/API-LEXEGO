// lexer/token.go

package lexer

// TokenType es un alias para string que mejora la legibilidad y seguridad del código.
// Representa el tipo de un token.
type TokenType string

// Token es la estructura que representa una unidad léxica del lenguaje.
type Token struct {
	Type    TokenType // El tipo del token (e.g., TOKEN_IDENTIFICADOR, TOKEN_NUMERO_ENTERO).
	Lexema  string    // El valor del token extraído del código fuente (e.g., "miVariable", "123").
	Linea   int       // La línea en el código fuente donde se encontró el token (para futuros reportes de errores).
	Columna int       // La columna en la línea donde comienza el token.
}

// --- Definición de Todos los Tipos de Token ---
// Estas constantes son una traducción directa de los estados finales de nuestro AFN.
const (
	// Tokens Especiales
	TOKEN_DESCONOCIDO TokenType = "DESCONOCIDO" // Token para un símbolo o secuencia no reconocida.
	TOKEN_EOF         TokenType = "EOF"         // Token especial para representar el Fin de Archivo (End of File).

	// Identificadores y Números (del AFN)
	TOKEN_ID     TokenType = "IDENTIFICADOR"
	TOKEN_NUM    TokenType = "NUMERO_ENTERO" // Corresponde a NUM
	TOKEN_INT_C  TokenType = "ENTERO_C"      // Corresponde a INT_C
	TOKEN_INT_L  TokenType = "ENTERO_LARGO"  // Corresponde a INT_L
	TOKEN_REAL   TokenType = "NUMERO_REAL"   // Corresponde a REAL
	TOKEN_SCIENT TokenType = "CIENTIFICO"    // Corresponde a SCIENT

	// Operadores (del AFN)
	// Nota: Los agrupamos como en el AFN, el lexer luego los diferenciará por su lexema.
	TOKEN_OP  TokenType = "OPERADOR"      // Para: +, -, *, /, %, (, )
	TOKEN_OP2 TokenType = "OPERADOR_COMP" // Para: ==, !=, <=, >=, :=, +=, etc.

	// Delimitadores y otros patrones (del AFN)
	TOKEN_DELIM  TokenType = "DELIMITADOR" // Para: >>>, <<<
	TOKEN_HEX    TokenType = "HEXADECIMAL" // Para: #[0-9A-Fa-f]{2,}
	TOKEN_STRING TokenType = "CADENA"      // Para: Q[0-9.-]+

	// NOTA: Mantenemos los nombres del AFN aunque parezcan cruzados.
	TOKEN_DINERO TokenType = "DINERO" // Para el patrón de tiempo HH:MM:SS
	TOKEN_HORA   TokenType = "HORA"   // Para el patrón de fecha \[DD/MM/YYYY\]
	TOKEN_FECHA  TokenType = "FECHA"  // Sin patrón definido en el AFN, pero se incluye por completitud.
)
