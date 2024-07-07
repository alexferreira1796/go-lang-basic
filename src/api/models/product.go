package models

import "api/src/api/db"

type Product struct {
	Id int
	Name string
	Description string
	Price float64
	Quantity int
} 

func SelectAllProducts() []Product {
	db := db.ConnectionDatabase() 

	selectAllProducts, err := db.Query("select * from products order by id desc")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	defer db.Close()
	return products
}

func SaveProduct(name, description string, price float64, quantity int) {
	db := db.ConnectionDatabase()

	insertProduct, err := db.Prepare("insert into products (name, description, price, quantity) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertProduct.Exec(name, description, price, quantity)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectionDatabase()

	deleteProduct, err := db.Prepare("delete from products where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)

	defer db.Close()
}

func GetProductSelected(id string) Product {
	db := db.ConnectionDatabase()

	productSelected, err := db.Query("select * from products where id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for productSelected.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productSelected.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}
	
	defer db.Close()

	return product
}

func UpdateProduct(id int, name string, description string, price float64, quantity int) {
	db := db.ConnectionDatabase()

	updatedProduct, err := db.Prepare("update products set name = $1, description = $2, price = $3, quantity = $4 where id = $5")
	if err != nil {
		panic(err.Error())
	}

	updatedProduct.Exec(name, description, price, quantity, id)

	defer db.Close()

}

// Slice => Array de produtos
// products := []Product{
// 	{
// 		Name: "Camiseta", 
// 		Description: "Bonita", 
// 		Price: 49.9, 
// 		Quantity: 3,
// 	},
// 	{
// 		"Tênis",
// 		"Confortável",
// 		329.9,
// 		5,
// 	},
// 	{
// 		"Fone",
// 		"Com Bluetooth",
// 		49.9,
// 		2,
// 	},
// }