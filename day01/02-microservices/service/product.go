package main

import(
	"errors"
)

type Product struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}

var products = []Product{
	{ID: 1, Name: "xiaomi 12", Price:99.99},
	{ID: 2, Name: "IPhone 13", Price: 88.88},
}

func (this Product) GetProducts(test string, prod *[]Product) error{
	*prod = products
	return nil
}

func (this Product) GetProductById(id int, prod *Product) error{
	for i := 0; i<len(products); i++{
		if id == products[i].ID {
			*prod = products[i]
			return nil
		}
	}
	return errors.New("查找不到该商品")
}
