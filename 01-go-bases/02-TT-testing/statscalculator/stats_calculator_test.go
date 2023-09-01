package statscalculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Los profesores de la universidad de Colombia, entraron al programa de análisis de datos  de Google,
	el cual premia a los mejores estadísticos de la región. Por ello los profesores nos solicitaron comprobar el correcto
	funcionamiento de nuestros cálculos estadísticos. Se solicita la siguiente tarea:
	Realizar test para calcular el mínimo de calificaciones.
	Realizar test para calcular el máximo de calificaciones.
	Realizar test para calcular el promedio de calificaciones.
*/

func TestMaximo(t *testing.T){
	t.Run("Calcular máximo de notas", func(t *testing.T){
		notas:= []float64{10, 20, 30, 40, 50}
		expectedResult := 50.0

		obtainResult := Maximo(notas...)


		assert.Equal(t, expectedResult, obtainResult)
		/*if(expectedResult != obtainResult){
			t.Fatal("Expected %f, obtain %f", expectedResult, obtainResult)
		}*/

	})
	t.Run("Calcular máximo de notas con error", func(t *testing.T){
		notas:= []float64{10, 20, 30, 40, 50, -10}
		expectedResult := 0.0
	
		obtainResult := Maximo(notas...)
	
		assert.Equal(t, expectedResult, obtainResult)
		/*if(expectedResult != obtainResult){
			t.Fatal("Expected %f, obtain %f", expectedResult, obtainResult)
		}*/
	
	})

}

func TestMinimo(t *testing.T){
	t.Run("Calcular mínimo de notas", func(t *testing.T){
		notas:= []float64{10, 20, 30, 40, 50}
		expectedResult := 10.0

		obtainResult := Minimo(notas...)
		
		assert.Equal(t, expectedResult, obtainResult)

	})

	t.Run("Calcular mínimo de notas con error", func(t *testing.T){
		notas:= []float64{10, 20, 30, 40, 50, -10}
		expectedResult := 0.0
	
		obtainResult := Minimo(notas...)
		
		assert.Equal(t, expectedResult, obtainResult)
	
	})

}

