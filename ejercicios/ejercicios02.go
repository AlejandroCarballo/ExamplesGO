package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Multiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)
	var num int
	var err error
	var text string

	for {
		fmt.Println("Ingrese un numero")
		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}
	for i := 1; i <= 10; i++ {
		text += fmt.Sprintf("%d x %d = %d \n", num, i, i*num)

	}
	return text

}
