package variables

import (
	"fmt"
	"strconv"
	"time"
)

var nombre string
var estado bool

var sueldo float32
var fecha time.Time

func RestoVariables() {

	nombre := "Alejandro"
	estado := true
	sueldo := 2500
	fecha = time.Now()

	fmt.Println(nombre)
	fmt.Println(estado)
	fmt.Println(sueldo)
	fmt.Println(fecha)

}

func ConviertoTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto

}
