package ejer_interfaces

import (
	"fmt"
	"github.com/AlejandroCarballo/ExamplesGO/interfaces"
)

func HumanosRespirando(h interfaces.Humano) {
	h.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", h.Sexo())
}
