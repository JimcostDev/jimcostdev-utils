# jimcostdev-utils

Este repositorio contiene utilidades en Go. Actualmente, incluye las siguientes funciones:

- **ToTitleCase:** Convierte una cadena de texto a formato *title case*, es decir, transforma la primera letra de cada palabra a mayúscula.
- **UniqueDigits:** Devuelve una cadena de dígitos únicos aleatorios. La longitud de la cadena será igual a `count`. Si `count` es mayor que 1, el primer dígito nunca será cero. La función retorna una cadena vacía si el valor de `count` es menor que 1 o mayor que 10.

## Características

- Ejemplo de pruebas unitarias utilizando *table-driven tests* para ambas funciones.
