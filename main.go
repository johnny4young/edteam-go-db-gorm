package main

import (
	"fmt"

	"github.com/johnny4young/edteam-go-db-gorm/model"
	"github.com/johnny4young/edteam-go-db-gorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	storage.DB().AutoMigrate(
		&model.Product{},
		&model.InvoiceHeader{},
		&model.InvoiceItem{},
	)

	// Create
	// product1 := model.Product{
	// 	Name:  "Go's Course",
	// 	Price: 120,
	// }

	// obs := "testing with GO"
	// product2 := model.Product{
	// 	Name:         "Go's Testing",
	// 	Price:        80,
	// 	Observations: &obs,
	// }
	// storage.DB().Create(&product1)
	// storage.DB().Create(&product2)

	// READ
	products := make([]model.Product, 0)
	storage.DB().Find(&products)
	for _, product := range products {
		fmt.Printf("%d - %s\n", product.ID, product.Name)
	}

}
