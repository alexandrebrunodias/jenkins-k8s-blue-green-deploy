package models

import (
	"strconv"

	"github.com/alexandrebrunodias/store/db"
)

// Product product dao
type Product struct {
	ID, Quantity      int
	Name, Description string
	Price             float64
}

// FindAllProducts find all products
func FindAllProducts() []Product {
	conn := db.Connect()

	result, err := conn.Query("select * from product order by ID")
	hasErrors(err)

	p := Product{}
	products := []Product{}

	for result.Next() {
		var ID, quantity int
		var name, description string
		var price float64

		err := result.Scan(&ID, &name, &description, &price, &quantity)
		hasErrors(err)

		p.ID = ID
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	defer conn.Close()
	return products
}

// CreateProduct create a product
func CreateProduct(name, description string, price float64, quantity int) {
	conn := db.Connect()

	insert, err := conn.Prepare("insert into product (name, description, price, quantity) values ($1, $2, $3, $4)")
	hasErrors(err)
	insert.Exec(name, description, price, quantity)

	defer conn.Close()
}

// DeleteProduct delete a product
func DeleteProduct(ID string) {
	conn := db.Connect()

	delete, err := conn.Prepare("delete from product where ID=$1")
	hasErrors(err)
	delete.Exec(ID)

	defer conn.Close()
}

// FindByID find product by ID
func FindByID(ID string) Product {
	conn := db.Connect()

	result, err := conn.Query("select * from product where ID=$1", ID)
	hasErrors(err)

	product := Product{}

	var quantity int
	var name, description string
	var price float64

	if result.Next() {
		err = result.Scan(&ID, &name, &description, &price, &quantity)
		hasErrors(err)

		product.ID, err = strconv.Atoi(ID)
		hasErrors(err)

		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}

	defer conn.Close()
	return product
}

// Update Update a product
func Update(ID, name, description, price, quantity string) {
	conn := db.Connect()

	update, err := conn.Prepare("update product set name=$2, description=$3, price=$4, quantity=$5 where ID=$1")
	hasErrors(err)

	_, err = update.Exec(ID, name, description, price, quantity)
	hasErrors(err)

	defer conn.Close()
}

func hasErrors(err error) {
	if err != nil {
		panic(err.Error())
	}
}
