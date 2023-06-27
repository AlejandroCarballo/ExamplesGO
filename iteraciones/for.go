package iteraciones

import "fmt"

func Iterar() {
	for i := 0; i < 10; i++ {
		if i == 6 {
			continue
		}

		fmt.Println("Hola For", i)

	}
}
