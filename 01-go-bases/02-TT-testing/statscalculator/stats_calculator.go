package statscalculator

import "fmt"

/*
Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de
los/as estudiantes de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de
enteros y devuelva el cálculo que se indicó en la función anterior.

*/

const (
	minimum="minimo"
	maximum="maximo"
	average="promedio"
)

func main(){
	var operacion string

	operacion = maximum
	funcion := orchestrator(operacion)
	fmt.Println(funcion(10, 20, 30, 40, 50, 60, 70, 80, 90, 100))

	operacion = minimum
	funcion = orchestrator(operacion)
	fmt.Println(funcion(10, 20, 30, 40, 50, 60, 70, 80, 90, 100))

	operacion = average
	funcion = orchestrator(operacion)
	fmt.Println(funcion(10, 20, 30, 40, 50, 60, 70, 80, 90, 100))
	

}

func orchestrator(tipoCalculo string, valores ...float64)(func (valores ...float64) float64){
	switch tipoCalculo{
	case minimum:
		return Minimo
	case maximum:
		return Maximo
	case average:
		return Promedio
	default:
		return nil
	}
}

func Maximo(valores ...float64)(max float64){
	max = valores[0]
	for _, valor := range valores{
		if(valor < 0){
			return 0
		}
		if(valor > max){
			max = valor
		}
	}
	return
}

func Minimo(valores ...float64)(min float64){
	min = valores[0]
	for _, valor := range valores{
		if(valor < 0){
			return 0
		}
		if(valor < min){
			min = valor
		}
	}
	return
}
func Promedio(valores ...float64)(prom float64){
	var suma float64 = 0
	for _, valor := range valores{
		if(valor < 0){
			return 0
		}
		suma += valor
	}
	prom = float64(suma) / float64(len(valores))
	return
}