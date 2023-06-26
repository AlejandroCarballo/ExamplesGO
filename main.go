package main

import (
	"fmt"
	"github.com/AlejandroCarballo/ExamplesGO/ejercicios"
)

func main() {
	/*variables.RestoVariables()
	estado, texto := variables.ConviertoTexto(1500)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "linux." || os == "OS x." {
		fmt.Println("Esto no es windows, es", os)

	} else {
		fmt.Println("Esto es windows")
	}

	switch os := runtime.GOOS; os {

	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es Darwin")

	default:
		fmt.Printf("%s \n", os)

	}*/

	numero, texto := ejercicios.Convertir("500")
	fmt.Println(numero)
	fmt.Println(texto)
}
