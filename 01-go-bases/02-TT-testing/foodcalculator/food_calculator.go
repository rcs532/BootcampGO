package foodcalculator

import "errors"

/*
Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.
Tienen los siguientes datos:
	Perro 10 kg de alimento.
	Gato 5 kg de alimento.
	Hamster 250 g de alimento.
	Tarántula 150 g de alimento.
Se solicita:
	Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado
	y que retorne una función y un mensaje (en caso que no exista el animal).
	Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
*/

func main(){
	var amount int
	var err error
	var f func(int) int

	f, err = animal("perro")
	if(err != nil){
		println(err)
	}else{
		amount = f(5)
		println(amount)
	}

	f, err = animal("gato")
	if(err != nil){
		println(err)
	}else{
		amount = f(5)
		println(amount)
	}

	f, err = animal("hamster")
	if(err != nil){
		println(err)
	}else{
		amount = f(5)
		println(amount)
	}
	

}


func animal(tipoAnimal string)(func (cantidad int) int, error){
	switch tipoAnimal{
	case "perro":
		return Perro, nil
	case "gato":
		return Gato, nil
	case "hamster":
		return Hamster, nil
	case "tarantula":
		return Tarantula, nil
	default:
		return nil, errors.New("No existe el animal")
	}
}

func Perro(cantidad int) int{
	return cantidad * 10
}
func Gato(cantidad int) int{
	return cantidad * 5
}
func Hamster(cantidad int) int{
	return cantidad * 250
}
func Tarantula(cantidad int) int{
	return cantidad * 150
}
