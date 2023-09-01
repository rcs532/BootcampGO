package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
	//arrange
	num:=1
	num2:=2
	expectedResult:=3

	//act
	obtainedResult:=Sumar(num,num2)

	//assert
	assert.Equal(t,expectedResult,obtainedResult) // si es igual no hay error
	/*if obtainedResult!=expectedResult{
		t.Fatal("Expected ",expectedResult," got ",obtainedResult)
	}*/

}

func TestRestar(t *testing.T) {
	result := Restar(2, 3)
	if result != -1 {
		t.Fatal("Expected -1, got ", result)
	}

	result = Restar(2, -3)
	if result != 5 {
		t.Fatal("Expected 5, got ", result)
	}

	result = Restar(-2, -3)
	if result != 1 {
		t.Fatal("Expected 1, got ", result)
	}
}

/*func TestDividir(t *testing.T) {
	//arrange
	num:=10.0
	num2:=2.0
	expectedResult:=5.0
	expectedErr:=false
	//act
	obtainedResult,obtainedErr:=Dividir(num,num2)

	//Assert
	if obtainedResult!=expectedResult{
		t.Fatal("Expected ",expectedResult," got ",obtainedResult)
	}
	if obtainedErr!=expectedErr{
		t.Fatal("Expected ",expectedErr," got ",obtainedErr)
	}
}*/

func TestDividir(t *testing.T) {
	t.Run("Dividir 10/2", func(t *testing.T) {
		//arrange
		num:=10.0
		num2:=0.0
		expectedResult:=0.0
		//expectedErr:=true
		//act
		obtainedResult,obtainedErr:=Dividir(num,num2)

		//Assert
		assert.False(t,obtainedErr)
		assert.Equal(t,expectedResult,obtainedResult)
		/*if obtainedResult!=expectedResult{
			t.Fatal("Expected ",expectedResult," got ",obtainedResult)
		}
		if obtainedErr!=expectedErr{
			t.Fatal("Expected ",expectedErr," got ",obtainedErr)
		}*/
	},
	)
	t.Run("Dividir 10/0", func(t *testing.T) {
		//arrange
		num:=10.0
		num2:=0.0
		expectedResult:=0.0
		//expectedErr:=true
		//act
		obtainedResult,obtainedErr:=Dividir(num,num2)

		//Assert
		assert.False(t,obtainedErr) // si es falso no hay error
		assert.Equal(t,expectedResult,obtainedResult) // si es igual no hay error
		/*if obtainedResult!=expectedResult{
			t.Fatal("Expected ",expectedResult," got ",obtainedResult)
		}
		if obtainedErr!=expectedErr{
			t.Fatal("Expected ",expectedErr," got ",obtainedErr)
		}*/
	},
	)
}

/*func TestDividirPorCero(t *testing.T) {
	//arrange
	num:=10.0
	num2:=0.0
	expectedResult:=0.0
	expectedErr:=true
	//act
	obtainedResult,obtainedErr:=Dividir(num,num2)

	//Assert
	if obtainedResult!=expectedResult{
		t.Fatal("Expected ",expectedResult," got ",obtainedResult)
	}
	if obtainedErr!=expectedErr{
		t.Fatal("Expected ",expectedErr," got ",obtainedErr)
	}
}*/

