package random

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// UniqueDigits devuelve una cadena de dígitos únicos aleatorios.
// La longitud de la cadena será igual a "count".
// Si count es mayor que 1, el primer dígito nunca será cero.
func UniqueDigits(count int) string {
	if count < 1 || count > 10 {
		return "" // No puede haber más de 10 dígitos únicos (0-9)
	}

	// Crear un generador de números aleatorios
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Crear un slice con todos los dígitos posibles
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Mezclar los dígitos aleatoriamente
	r.Shuffle(len(digits), func(i, j int) {
		digits[i], digits[j] = digits[j], digits[i]
	})

	// Si es un número de múltiples dígitos, asegurarse que no comience con cero
	if count > 1 && digits[0] == 0 {
		// Encontrar el primer dígito no cero
		for i := 1; i < len(digits); i++ {
			if digits[i] != 0 {
				// Intercambiar
				digits[0], digits[i] = digits[i], digits[0]
				break
			}
		}
	}

	// Construir la cadena de resultado
	var result strings.Builder
	for i := 0; i < count; i++ {
		result.WriteString(strconv.Itoa(digits[i]))
	}

	return result.String()
}
