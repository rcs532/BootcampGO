package taxcalculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	La  empresa de chocolates que anteriormente necesitaba calcular el impuesto de sus empleados, al momento de depositar el sueldo de
	los mismos ahora nos solicitó validar que los cálculos de estos impuestos están correctos. Para esto nos encargaron el trabajo de realizar
	los test correspondientes para:
	Calcular el impuesto en caso de que el empleado gane por debajo de $50.000.
	Calcular el impuesto en caso de que el empleado gane por encima de $50.000.
	Calcular el impuesto en caso de que el empleado gane por encima de $150.000.
*/

func TestCalculadoraDeImpuestos(t *testing.T){
	t.Run("Calcular impuesto cuando el salario es menor a 50000", func(t *testing.T){
	salario := 70000.0
	expectedResult := 11900.0

	obtainedResult := CalcularImpuesto(salario)

	assert.Equal(t, expectedResult, obtainedResult, "El impuesto calculado no es el esperado")
	})
	t.Run("Calcular impuesto cuando el salario es mayor a 50000 y menor a 150000", func(t *testing.T){
		salario := 70000.0
		expectedResult := 11900.0
	
		obtainedResult := CalcularImpuesto(salario)
	
		assert.Equal(t, expectedResult, obtainedResult, "El impuesto calculado no es el esperado")
	})
	t.Run("Calcular impuesto cuando el salario es mayor a 150000", func(t *testing.T){
		salario := 70000.0
		expectedResult := 11900.0
	
		obtainedResult := CalcularImpuesto(salario)
	
		assert.Equal(t, expectedResult, obtainedResult, "El impuesto calculado no es el esperado")
	})
}
