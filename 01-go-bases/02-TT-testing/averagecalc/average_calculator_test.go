package averagecalc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	El colegio informó que las operaciones para calcular el promedio no se están realizando correctamente,
	por lo que ahora nos corresponde realizar los test correspondientes:
	Calcular el promedio de las notas de los alumnos.
*/

func TestCalcularPromedio(t *testing.T){
	t.Run("Calcular promedio de notas", func(t *testing.T){
	notas:= []int{10, 20, 30, 40, 50}
	expectedResult := 30.0

	obtainResult, _ := CalcularPromedios(notas...)

	assert.Equal(t, expectedResult, obtainResult)

	})
	t.Run("Calcular promedio de notas con error", func(t *testing.T){
		notas:= []int{10, 20, 30, 40, 50, -10}
		expectedResult := 0.0
	
		obtainResult, _ := CalcularPromedios(notas...)
	
		assert.Equal(t, expectedResult, obtainResult)
	
	})

}