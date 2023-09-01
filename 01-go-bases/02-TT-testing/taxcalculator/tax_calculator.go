package taxcalculator

import "fmt"

/*
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más
de $150.000 se le descontará además un 10 % (27% en total).

*/


func main(){
	var salario float64 = 70000
	var impuesto float64 = CalcularImpuesto(salario)
	fmt.Println(impuesto)


}

func CalcularImpuesto(salario float64)float64{
	var impuesto float64 = 0
	if salario > 50000 && salario < 150000{
		impuesto = salario * 0.17
	}else if salario > 150000{
		impuesto = salario * 0.27
	}
	return impuesto
}