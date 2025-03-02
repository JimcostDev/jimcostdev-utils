package random

import (
	"testing"
	"unicode"
)

func TestUniqueDigits(t *testing.T) {
	tests := []struct {
		name  string
		count int
		valid bool // indica si se espera un resultado válido o ""
	}{
		{
			name:  "Cuenta 0 (inválido)",
			count: 0,
			valid: false,
		},
		{
			name:  "Cuenta negativa (inválido)",
			count: -1,
			valid: false,
		},
		{
			name:  "Cuenta mayor que 10 (inválido)",
			count: 11,
			valid: false,
		},
		{
			name:  "Un dígito (válido)",
			count: 1,
			valid: true,
		},
		{
			name:  "Dos dígitos (válido, primer dígito no es cero)",
			count: 2,
			valid: true,
		},
		{
			name:  "Diez dígitos (válido, todos los dígitos únicos)",
			count: 10,
			valid: true,
		},
	}

	for _, tc := range tests {
		tc := tc // capturar variable para cada subtest
		t.Run(tc.name, func(t *testing.T) {
			result := UniqueDigits(tc.count)
			// Para counts inválidos, se espera cadena vacía.
			if !tc.valid {
				if result != "" {
					t.Errorf("Para count=%d se esperaba una cadena vacía, pero se obtuvo %q", tc.count, result)
				}
				return
			}

			// Para counts válidos, verificamos la longitud.
			if len(result) != tc.count {
				t.Errorf("Para count=%d se esperaba longitud %d, pero se obtuvo %d en %q", tc.count, tc.count, len(result), result)
			}

			// Verificar que cada carácter sea un dígito y no se repita.
			seen := make(map[rune]bool)
			for i, r := range result {
				if !unicode.IsDigit(r) {
					t.Errorf("El carácter %q en la posición %d no es un dígito", r, i)
				}
				if seen[r] {
					t.Errorf("El dígito %q se repite en la cadena %q", r, result)
				}
				seen[r] = true
			}

			// Si count > 1, el primer dígito no debe ser '0'
			if tc.count > 1 && result[0] == '0' {
				t.Errorf("Para count=%d, el primer dígito no debe ser '0', pero se obtuvo %q", tc.count, result)
			}
		})
	}
}
