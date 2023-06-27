package arreglos_slices

import "fmt"

var tablaSlice []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{2, 5, 4, 8, 9, 2, 46, 8}

func MuestroSlices() {
	fmt.Println(tablaSlice)

	porcion1 := arreglo[3:]  //desde la posicion 3
	porcion2 := arreglo[:5]  //desde la posicion 0 hasta la 5
	porcion3 := arreglo[6:7] // desde la posicion 6 hasta la 7

	fmt.Println(porcion1)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)

	fmt.Printf("Largo %d, capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
		fmt.Printf("\nLargo %d, capacidad %d", len(nums), cap(nums))

	}
}
