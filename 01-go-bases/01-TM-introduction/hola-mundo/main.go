package main

import "fmt" //nos permite utilizar la entrada y salida de datos


func main(){
	fmt.Println("Hola Mundo!") //comentarios

	var username string = "Juan" //declaracion de variables

	//make a pointer
	var pointer *string = &username

	var userNameModified string = *pointer + " Pablo"

	

	fmt.Println(userNameModified)

	/*
	Comentario multilinea
	*/
}