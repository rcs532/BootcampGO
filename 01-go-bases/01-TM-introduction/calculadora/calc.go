package calculadora

func Sumar(a int, b int) int {
	return a + b
}

func Restar(a int, b int) int {
	return a - b
}

func Dividir(a float64, b float64) (result float64, err bool) {
	if b == 0 {
		err = true
		return
	}
	result =a / b
	return
}
