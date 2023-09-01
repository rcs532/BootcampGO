package main

import "structs/employee"

func main(){
	person := employee.Person{
		ID: 1,
		Name: "John Doe",
		DateOfBirth: "1990-01-01",
	}
	employee:= employee.Employee{
			ID: 12,
			Position: "Software Engineer",
			Person: person,
	}
	employee.PrintEmployee()
}