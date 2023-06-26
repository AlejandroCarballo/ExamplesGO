package main

import (
	"fmt"
	"github.com/AlejandroCarballo/ExamplesGO/variables"
)

func main() {
	variables.RestoVariables()
	estado, texto := variables.ConviertoTexto(1500)
	fmt.Println(estado)
	fmt.Println(texto)
}
