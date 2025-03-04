package main

import (
	"fmt"

	"github.com/JimcostDev/jimcostdev-utils/random"
	"github.com/JimcostDev/jimcostdev-utils/strutils"
)

func main() {
	texto := "hola mundo desde go"
	resultado := strutils.ToTitleCase(texto)
	fmt.Println(resultado) // Imprime: Hola Mundo Desde Go

	secreto := random.UniqueDigits(4)
	fmt.Println(secreto)
}
