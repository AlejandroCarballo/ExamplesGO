package main

import "github.com/AlejandroCarballo/ExamplesGO/middleware"

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

	}

	numero, texto := ejercicios.Convertir("500")
	fmt.Println(numero)
	fmt.Println(texto)*/

	//teclado.IngresoNumeros()
	//iteraciones.Iterar()

	//fmt.Println(ejercicios.Multiplicar())
	//files.GrabaTable()

	//files.SumaTabla()
	//files.LeoArchivo()
	//funciones.LlamarClosure()

	//funciones.Exponencia(2)
	//maps.MostrarMapas()
	//users.AltaUsuario()
	/*Ale := new(modelos.Hombre)
	ejer_interfaces.HumanosRespirando(Ale)

	Melisa := new(modelos.Mujer)
	ejer_interfaces.HumanosRespirando(Melisa)*/

	//defer_panic.EjemploPanic()

	/*canal1 := make(chan bool)
	go goroutines.MiNombreLeentooo("Alejandro Carballo", canal1)

	<-canal1

	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)*/
	//webserver.MiWebServer()
	middleware.MiMiddleware()

}
