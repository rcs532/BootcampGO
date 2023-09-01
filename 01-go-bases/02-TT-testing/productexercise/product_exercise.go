package main

import "fmt"

type Product struct {
	ID int
	Name string
	Price float64
	Category string
}

// slice global de products llamado Products
var Products []Product = []Product{
	{
		 1,
		"Coca Cola",
		 20.5,
		"Drink",
	},
	{
		 2,
		"Sprite",
		 16.5,
		"Drink",
	},
	{
		 3,
		"Manzana",
		 10.5,
		"Drink",
	},
}

func (p Product) GetAll()  {
	for _, product := range Products {
		fmt.Println(product)
	}

}

func (p Product) GetByID(id int) Product {
	var producto Product
	for _, product := range Products {
		if product.ID == id {
			producto = product
		}
	}
	return producto
}

func (p Product) Save() {
	Products = append(Products, p)
}

func main() {
	var product Product
	product.GetAll()

	product = product.GetByID(1)
	fmt.Println(product)

	product = Product{
		ID:       4,
		Name:     "Pepsi",
		Price:    20.5,
		Category: "Drink",

	}
	product.Save()
	product.GetAll()
}
