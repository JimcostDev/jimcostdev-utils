package strutils

import "testing"

// TestToTitleCase es una función de prueba para ToTitleCase
func TestToTitleCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Cadena vacía",
			input:    "",
			expected: "",
		},
		{
			name:     "Una sola palabra",
			input:    "hola",
			expected: "Hola",
		},
		{
			name:     "Varias palabras",
			input:    "hola mundo",
			expected: "Hola Mundo",
		},
		{
			name:     "Con espacios extra",
			input:    "   hola    mundo  ",
			expected: "Hola Mundo", // strings.Fields elimina espacios redundantes
		},
		{
			name:     "Ya en formato título",
			input:    "Hola Mundo",
			expected: "Hola Mundo",
		},
	}

	for _, tc := range tests {
		tc := tc // capturar variable para cada subtest
		t.Run(tc.name, func(t *testing.T) {
			result := ToTitleCase(tc.input)
			if result != tc.expected {
				t.Errorf("Para %q se esperaba %q, pero se obtuvo %q", tc.input, tc.expected, result)
			}
		})
	}
}
