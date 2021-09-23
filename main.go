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

	//ALL
	products := make([]model.Product, 0)
	storage.DB().Find(&products)
	for _, product := range products {
		fmt.Printf("%d - %s\n", product.ID, product.Name)
	}

	//get one row
	myProduct := model.Product{}
	storage.DB().First(&myProduct, 1)
	fmt.Println(myProduct)

	//UPDATE
	myProduct2 := model.Product{}
	myProduct2.ID = 2
	storage.DB().Model(&myProduct2).Updates(
		model.Product{Name: "Java's Course", Price: 55},
	)

	// Delete Soft
	myProduct3 := model.Product{}
	myProduct3.ID = 4

	storage.DB().Delete(&myProduct3)

	// Delete Permanent
	myProduct4 := model.Product{}
	myProduct4.ID = 4

	storage.DB().Unscoped().Delete(&myProduct4)

}
