// lexer/lexer.go

package lexer

import (
	"unicode"
)

// AnalizadorLexico define la estructura que maneja el estado del análisis.
type AnalizadorLexico struct {
	fuente   []rune
	posicion int
	linea    int
	columna  int
	caracter rune
}

// NuevoAnalizador crea y devuelve un nuevo analizador léxico listo para usarse.
func NuevoAnalizador(fuente string) *AnalizadorLexico {
	a := &AnalizadorLexico{fuente: []rune(fuente), linea: 1, columna: 1}
	a.leerCaracter()
	return a
}

// SiguienteToken es el método principal que consume la fuente y devuelve el siguiente token.
// Actúa como el "director de orquesta", delegando tareas a los reconocedores.
func (a *AnalizadorLexico) SiguienteToken() Token {
	for {
		a.saltarEspaciosYComentarios()
		lineaActual, columnaActual := a.linea, a.columna

		switch a.caracter {
		case 0:
			return Token{Type: TOKEN_EOF, Lexema: "", Linea: lineaActual, Columna: columnaActual}
		case '#':
			return a.reconocerHexadecimal()
		case 'Q':
			return a.reconocerCadenaString()
		case '[':
			return a.reconocerHora()
		case '&':
			if a.peek() == '&' {
				a.leerCaracter()
				return a.nuevoToken(TOKEN_OP2, lineaActual, columnaActual)
			}
		case '|':
			if a.peek() == '|' {
				a.leerCaracter()
				return a.nuevoToken(TOKEN_OP2, lineaActual, columnaActual)
			}
		case '(':
			return a.nuevoToken(TOKEN_OP, lineaActual, columnaActual)
		case ')':
			return a.nuevoToken(TOKEN_OP, lineaActual, columnaActual)
		case '+':
			if a.peek() == '=' {
				a.leerCaracter()
				return a.nuevoToken(TOKEN_OP2, lineaActual, columnaActual)
			}
			return a.nuevoToken(TOKEN_OP, lineaActual, columnaActual)
		case '-':
			if a.peek() == '=' {
				a.leerCaracter()
				return a.nuevoToken(TOKEN_OP2, lineaActual, columnaActual)
			}
			return a.nuevoToken(TOKEN_OP, lineaActual, columnaActual)
		case '*':
			if a.peek() == '=' {
				a.leerCaracter()
				return a.nuevoToken(TOKEN_OP2, lineaActual, columnaActual)
			}
			return a.nuevoToken(TOKEN_OP, lineaActual, columnaActual)
		case '/':
			if a.peek() == '=' {
				a.leerCaracter()
				return a.nuevoToken(TOKEN_OP2, lineaActual, columnaActual)
			}
			return a.nuevoToken(TOKEN_OP, lineaActual, columnaActual)
		case '%':
			if a.peek() == '=' {
				a.leerCaracter()
				return a.nuevoToken(TOKEN_OP2, lineaActual, columnaActual)
			}
			return a.nuevoToken(TOKEN_OP, lineaActual, columnaActual)
		case '=':
			if a.peek() == '=' {
				a.leerCaracter()
				return a.nuevoToken(TOKEN_OP2, lineaActual, columnaActual)
			}
			return a.nuevoToken(TOKEN_DESCONOCIDO, lineaActual, columnaActual)
		case ':':
			if a.peek() == '=' {
				a.leerCaracter()
				return a.nuevoToken(TOKEN_OP2, lineaActual, columnaActual)
			}
			return a.nuevoToken(TOKEN_DESCONOCIDO, lineaActual, columnaActual)
		case '!':
			if a.peek() == '=' {
				a.leerCaracter()
				return a.nuevoToken(TOKEN_OP2, lineaActual, columnaActual)
			}
			return a.nuevoToken(TOKEN_OP2, lineaActual, columnaActual)
		case '<':
			if a.peek() == '<' && a.peekMasAdelante(2) == '<' {
				a.leerCaracter()
				a.leerCaracter()
				return a.nuevoToken(TOKEN_DELIM, lineaActual, columnaActual)
			}
			if a.peek() == '=' {
				a.leerCaracter()
				return a.nuevoToken(TOKEN_OP2, lineaActual, columnaActual)
			}
			return a.nuevoToken(TOKEN_OP2, lineaActual, columnaActual)
		case '>':
			if a.peek() == '>' && a.peekMasAdelante(2) == '>' {
				a.leerCaracter()
				a.leerCaracter()
				return a.nuevoToken(TOKEN_DELIM, lineaActual, columnaActual)
			}
			if a.peek() == '=' {
				a.leerCaracter()
				return a.nuevoToken(TOKEN_OP2, lineaActual, columnaActual)
			}
			return a.nuevoToken(TOKEN_OP2, lineaActual, columnaActual)
		default:
			if unicode.IsLetter(a.caracter) || a.caracter == '_' {
				return Token{Type: TOKEN_ID, Lexema: a.reconocerIdentificador(), Linea: lineaActual, Columna: columnaActual}
			} else if unicode.IsDigit(a.caracter) {
				return a.reconocerNumero()
			} else {
				return a.nuevoToken(TOKEN_DESCONOCIDO, lineaActual, columnaActual)
			}
		}
		return a.nuevoToken(TOKEN_DESCONOCIDO, lineaActual, columnaActual)
	}
}

// --- Funciones de Ayuda de Bajo Nivel ---

// leerCaracter avanza la posición en la fuente.
func (a *AnalizadorLexico) leerCaracter() {
	if a.posicion >= len(a.fuente) {
		a.caracter = 0
	} else {
		a.caracter = a.fuente[a.posicion]
	}
	a.posicion++
	a.columna++
}

// peek espía el siguiente carácter sin consumir.
func (a *AnalizadorLexico) peek() rune {
	if a.posicion >= len(a.fuente) {
		return 0
	}
	return a.fuente[a.posicion]
}

// peekMasAdelante espía N caracteres más adelante.
func (a *AnalizadorLexico) peekMasAdelante(offset int) rune {
	if a.posicion+offset-1 >= len(a.fuente) {
		return 0
	}
	return a.fuente[a.posicion+offset-1]
}

// nuevoToken crea un nuevo token a partir del carácter actual y avanza.
// BUG CORREGIDO: El lexema para TOKEN_DELIM ahora se calcula correctamente.
func (a *AnalizadorLexico) nuevoToken(tipo TokenType, linea, columna int) Token {
	posicionInicial := a.posicion - 1
	var lexema string

	switch tipo {
	case TOKEN_DELIM:
		lexema = string(a.fuente[posicionInicial-2 : posicionInicial+1])
	case TOKEN_OP2:
		lexema = string(a.fuente[posicionInicial-1 : posicionInicial+1])
	default: // TOKEN_OP, TOKEN_DESCONOCIDO, etc.
		lexema = string(a.fuente[posicionInicial])
	}

	a.leerCaracter()
	return Token{Type: tipo, Lexema: lexema, Linea: linea, Columna: columna}
}
