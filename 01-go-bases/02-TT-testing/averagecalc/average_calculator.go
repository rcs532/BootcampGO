package averagecalc

import (
	"errors"
	"fmt"
)

/*
Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio.
No se pueden introducir notas negativas.

*/

func main(){
	var promedio float64
	var err error
	promedio, err = CalcularPromedios(10, 20, 30, 40, 50, -10)
	if(err != nil){
		fmt.Println(err)

	}else{
		fmt.Println(promedio, err)
	}

	

}

func CalcularPromedios(valores ...int)(promedio float64, err error){
	var suma int = 0
	err = nil
	for _, valor := range valores{
		if(valor < 0){
			err = errors.New("No se puede introducir notas negativas")
			promedio = 0 // asignar cero al promedio cuando hay un error
			return
		}
		suma += valor
	}
	promedio = float64(suma) / float64(len(valores))
	return 

}
