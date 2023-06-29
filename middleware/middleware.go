package middleware

import (
	"fmt"
)

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")

	res := operacionesMidd(sumar)(2, 3)
	fmt.Println(res)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {

	return func(a int, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)

	}

}
