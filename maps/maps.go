package maps

import "fmt"

func MostrarMapas() {

	paises := make(map[string]string, 5)
	fmt.Println(paises)

	paises["Argentina"] = "Buenos Aires"
	paises["Mexico"] = "DF"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barcelona": 39,
		"Boca":      41,
		"River":     10,
		"Belgrano":  110,
	}
	fmt.Println(campeonato)

	/*for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un porcentaje de %d\n", equipo, puntaje)
	}*/

	delete(campeonato, "Boca")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["River"]
	fmt.Printf("El puntaje es %d, y el equipo existe = %t \n", puntaje, existe)

}
