// lexer/recognizers.go

package lexer

import (
	"strings"
	"unicode"
)

// saltarEspaciosYComentarios consume todos los espacios en blanco, saltos de línea y comentarios.
func (a *AnalizadorLexico) saltarEspaciosYComentarios() {
	for {
		if a.caracter == ' ' || a.caracter == '\t' || a.caracter == '\n' || a.caracter == '\r' {
			if a.caracter == '\n' {
				a.linea++
				a.columna = 0
			}
			a.leerCaracter()
		} else if a.caracter == '-' && a.peek() == '-' {
			for a.caracter != '\n' && a.caracter != 0 {
				a.leerCaracter()
			}
		} else {
			break
		}
	}
}

// reconocerIdentificador consume los caracteres de un ID y devuelve el lexema.
func (a *AnalizadorLexico) reconocerIdentificador() string {
	posicionInicial := a.posicion - 1
	for unicode.IsLetter(a.caracter) || unicode.IsDigit(a.caracter) || a.caracter == '_' {
		a.leerCaracter()
	}
	return string(a.fuente[posicionInicial : a.posicion-1])
}

// reconocerHora consume y valida el formato de fecha [DD/MM/YYYY].
func (a *AnalizadorLexico) reconocerHora() Token {
	lineaActual, columnaActual := a.linea, a.columna
	posicionInicial := a.posicion - 1

	if a.posicion+10 > len(a.fuente) {
		a.leerCaracter()
		return Token{Type: TOKEN_DESCONOCIDO, Lexema: string(a.fuente[posicionInicial]), Linea: lineaActual, Columna: columnaActual}
	}

	subcadena := string(a.fuente[posicionInicial : a.posicion+11])

	if subcadena[0] == '[' && unicode.IsDigit(rune(subcadena[1])) && unicode.IsDigit(rune(subcadena[2])) &&
		subcadena[3] == '/' && unicode.IsDigit(rune(subcadena[4])) && unicode.IsDigit(rune(subcadena[5])) &&
		subcadena[6] == '/' && unicode.IsDigit(rune(subcadena[7])) && unicode.IsDigit(rune(subcadena[8])) &&
		unicode.IsDigit(rune(subcadena[9])) && unicode.IsDigit(rune(subcadena[10])) && subcadena[11] == ']' {

		for i := 0; i < 12; i++ {
			a.leerCaracter()
		}
		return Token{Type: TOKEN_HORA, Lexema: subcadena, Linea: lineaActual, Columna: columnaActual}
	}

	a.leerCaracter()
	return Token{Type: TOKEN_DESCONOCIDO, Lexema: string(a.fuente[posicionInicial]), Linea: lineaActual, Columna: columnaActual}
}

// reconocerHexadecimal consume y valida un número hexadecimal.
func (a *AnalizadorLexico) reconocerHexadecimal() Token {
	lineaActual, columnaActual := a.linea, a.columna
	posicionInicial := a.posicion - 1
	a.leerCaracter()

	for unicode.IsDigit(a.caracter) || (unicode.ToLower(a.caracter) >= 'a' && unicode.ToLower(a.caracter) <= 'f') {
		a.leerCaracter()
	}

	lexema := string(a.fuente[posicionInicial : a.posicion-1])
	if len(lexema) < 3 {
		return Token{Type: TOKEN_DESCONOCIDO, Lexema: lexema, Linea: lineaActual, Columna: columnaActual}
	}
	return Token{Type: TOKEN_HEX, Lexema: lexema, Linea: lineaActual, Columna: columnaActual}
}

// reconocerCadenaString consume y valida una cadena que empieza con 'Q'.
func (a *AnalizadorLexico) reconocerCadenaString() Token {
	lineaActual, columnaActual := a.linea, a.columna
	posicionInicial := a.posicion - 1
	a.leerCaracter()

	for unicode.IsDigit(a.caracter) || a.caracter == '.' || a.caracter == '-' {
		a.leerCaracter()
	}

	lexema := string(a.fuente[posicionInicial : a.posicion-1])
	if len(lexema) < 2 {
		return Token{Type: TOKEN_DESCONOCIDO, Lexema: lexema, Linea: lineaActual, Columna: columnaActual}
	}
	return Token{Type: TOKEN_STRING, Lexema: lexema, Linea: lineaActual, Columna: columnaActual}
}

// reconocerNumero consume y valida todos los tipos de números (entero, real, científico, etc.).
func (a *AnalizadorLexico) reconocerNumero() Token {
	lineaActual, columnaActual := a.linea, a.columna
	posicionInicial := a.posicion - 1
	tipo := TOKEN_NUM

	for unicode.IsDigit(a.caracter) {
		a.leerCaracter()
	}
	if a.caracter == '.' {
		tipo = TOKEN_REAL
		a.leerCaracter()
		if !unicode.IsDigit(a.caracter) {
			return Token{Type: TOKEN_DESCONOCIDO, Lexema: string(a.fuente[posicionInicial : a.posicion-1]), Linea: lineaActual, Columna: columnaActual}
		}
		for unicode.IsDigit(a.caracter) {
			a.leerCaracter()
		}
	}
	if unicode.ToUpper(a.caracter) == 'E' {
		tipo = TOKEN_SCIENT
		a.leerCaracter()
		if a.caracter == '+' || a.caracter == '-' {
			a.leerCaracter()
		}
		if !unicode.IsDigit(a.caracter) {
			return Token{Type: TOKEN_DESCONOCIDO, Lexema: string(a.fuente[posicionInicial : a.posicion-1]), Linea: lineaActual, Columna: columnaActual}
		}
		for unicode.IsDigit(a.caracter) {
			a.leerCaracter()
		}
	}
	if a.caracter == '_' && tipo == TOKEN_NUM {
		a.leerCaracter()
		if a.caracter == 'c' {
			a.leerCaracter()
			tipo = TOKEN_INT_C
		} else if strings.HasPrefix(string(a.fuente[a.posicion-1:]), "largo") {
			for i := 0; i < 5; i++ {
				a.leerCaracter()
			}
			tipo = TOKEN_INT_L
		}
	}

	lexema := string(a.fuente[posicionInicial : a.posicion-1])
	return Token{Type: tipo, Lexema: lexema, Linea: lineaActual, Columna: columnaActual}
}
