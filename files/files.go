package files

import (
	"bufio"
	"fmt"
	"github.com/AlejandroCarballo/ExamplesGO/ejercicios"
	"os"
)

var fileName string = "./files/txt/tabla.txt"

func GrabaTable() {
	var text string = ejercicios.Multiplicar()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}

	fmt.Fprintln(archivo, text)
	archivo.Close()
}

func SumaTabla() {
	var text string = ejercicios.Multiplicar()
	if !Append(fileName, text) {
		fmt.Println("Error al concatenar contenido")
	}

}

func Append(file string, text string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append", err.Error())
		return false
	}
	_, err = arch.WriteString(text)
	if err != nil {
		fmt.Println("Error durante el WriteString", err.Error())
		return false
	}
	arch.Close()
	return true

}

func LeoArchivo() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println(">" + registro)

	}
	archivo.Close()
}
