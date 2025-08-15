# jimcostdev-utils

Utilidades Go para manipulación de cadenas y generación de números aleatorios únicos.

## Descripción

Paquete de utilidades en Go que incluye funciones para:

- Formatear cadenas en estilo título (`strutils.TitleCase`)
- Generar números aleatorios de `n` dígitos sin repetir dígitos (`random.RandomNDigits`)

Ideal para proyectos que requieren manipulación avanzada de texto y generación de identificadores numéricos únicos.

## Instalación

```bash
go get github.com/JimcostDev/jimcostdev-utils@latest
```

## Uso

```go
import "github.com/JimcostDev/jimcostdev-utils/strutils"
import "github.com/JimcostDev/jimcostdev-utils/random"
```

## Ejemplo

```go
title := strutils.TitleCase("hola mundo")
num := random.RandomNDigits(5)
```

## Licencia

MIT
