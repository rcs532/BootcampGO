package employee

import "fmt"

//Person is a struct that represents a person
type Person struct {
	ID int
	Name string
	DateOfBirth string
}

type Employee struct {
	Person
	ID int
	Position string
}

func (e Employee) PrintEmployee(){
	fmt.Println("ID: ", e.ID)
	fmt.Println("Name: ", e.Name)
	fmt.Println("Date of Birth: ", e.DateOfBirth)
	fmt.Println("Position: ", e.Position)
}

