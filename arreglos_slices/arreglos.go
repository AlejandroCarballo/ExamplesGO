package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{10, 0, 59}

var matriz [20][30]int

func MuestroArreglo() {

	tabla[7] = 12
	tabla[2] = 2
	tabla2 := [10]string{"Pablo", "Juan", "Ale"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])

	}

	matriz[5][8] = 100
	fmt.Println(matriz)

}
