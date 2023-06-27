package funciones

import "fmt"

func tabla(valor int) func() int {
	num := valor
	secuencia := 0
	return func() int {
		secuencia++
		return num * secuencia
	}
}

func LlamarClosure() {
	tabladel := 2
	MiTable := tabla(tabladel)
	for i := 0; i <= 10; i++ {
		fmt.Println(MiTable())

	}
}
